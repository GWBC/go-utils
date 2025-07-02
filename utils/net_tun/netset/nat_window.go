//go:build windows

package netset

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/GWBC/go-utils/utils"
	"github.com/lysShub/divert-go"
	"golang.org/x/sys/windows"
	"gvisor.dev/gvisor/pkg/tcpip"
	"gvisor.dev/gvisor/pkg/tcpip/header"
)

type AddrInfo struct {
	Addr     tcpip.Address
	Port     uint16
	lastTime int64
}

var wg sync.WaitGroup
var natsLock sync.RWMutex
var nats []net.IPNet
var restart chan bool
var ctx context.Context
var cancel context.CancelFunc

var portMapLock sync.RWMutex
var portMap map[uint16]*AddrInfo

var _ = divert.MustLoad(divert.DLL)

func StartForward() {
	ctx, cancel = context.WithCancel(context.Background())
	restart = make(chan bool)
	nats = []net.IPNet{}
	portMap = map[uint16]*AddrInfo{}

	wg.Add(1)
	go forward()
}

func StopForward() {
	select {
	case <-ctx.Done():
		return
	default:
		cancel()
		close(restart)
	}
}

func DelNatMasquerade(netAddr net.IPNet) {
	select {
	case <-ctx.Done():
		return
	default:
		natsLock.Lock()
		for i, nat := range nats {
			if netAddr.String() == nat.String() {
				copy(nats[i:], nats[i+1:])
			}
		}
		natsLock.Unlock()
	}

	restart <- true
}

func AddNatMasquerade(netAddr net.IPNet) error {
	select {
	case <-ctx.Done():
		return errors.New("nats is stop")
	default:
		natsLock.Lock()
		nats = append(nats, utils.CopyIPNet(netAddr))
		natsLock.Unlock()
	}

	restart <- true
	return nil
}

func forward() {
	var dev *divert.Handle

	closeDev := func() {
		if dev != nil {
			dev.Close()
			dev = nil
		}
	}

	genFilter := func() string {
		return "!loopback"
	}

	restartDev := func(filter string) {
		closeDev()

		d, err := divert.Open(filter, divert.Network, 0, 0)
		if err != nil {
			time.Sleep(1 * time.Second)
			restart <- true
			return
		}

		var addr divert.Address
		var data = make([]byte, 64*1024)

		wg.Add(1)
		go func() {
			defer wg.Done()

			isNeedNat := func(ip *tcpip.Address) bool {
				netip := net.IP(ip.AsSlice())
				natsLock.RLock()
				defer natsLock.RUnlock()
				for _, nat := range nats {
					if nat.Contains(netip) {
						return true
					}
				}

				return false
			}

			for {
				n, err := d.Recv(data, &addr)
				if err != nil {
					if errors.Is(err, windows.ERROR_INSUFFICIENT_BUFFER) {
						continue
					}

					return
				}

				//IPV6
				if addr.IPv6() {
					d.Send(data, &addr)
					continue
				}

				//IPV4
				if n >= header.IPv4MinimumSize+header.TCPMinimumSize {
					iphdr := header.IPv4(data[:n])

					tcphdr := header.TCP(iphdr[iphdr.HeaderLength():])

					iphdr.SetSourceAddress()
					fmt.Printf("%s:%d --> %s:%d \n",
						iphdr.SourceAddress().String(),
						tcphdr.SourcePort(),
						iphdr.DestinationAddress().String(),
						tcphdr.DestinationPort(),
					)

					d.Send(data, &addr)
				}
			}
		}()
	}

	defer func() {
		closeDev()
		wg.Done()
	}()

	for {
		select {
		case <-restart:
			isRestart := func() bool {
				natsLock.RLock()
				defer natsLock.RUnlock()
				return len(nats) != 0
			}()

			if !isRestart {
				continue
			}

			restartDev(genFilter())
		case <-ctx.Done():
			return
		}
	}
}

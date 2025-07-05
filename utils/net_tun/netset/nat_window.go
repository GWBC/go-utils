//go:build windows

package netset

import (
	"context"
	"encoding/binary"
	"errors"
	"net"
	"os/exec"
	"regexp"
	"sync"
	"time"

	"github.com/GWBC/go-utils/utils"
	expiremap "github.com/GWBC/go-utils/utils/expire_map"
	"github.com/lysShub/divert-go"
	"gvisor.dev/gvisor/pkg/tcpip"
	"gvisor.dev/gvisor/pkg/tcpip/checksum"
	"gvisor.dev/gvisor/pkg/tcpip/header"
)

const NAT_CACHE_TIMEOUT = 5 * time.Minute

var wg sync.WaitGroup
var natsLock sync.RWMutex
var nats []net.IPNet
var restart chan bool
var ctx context.Context
var cancel context.CancelFunc

var defaultAddr net.IP
var loaclIPMap map[uint32]bool

var natMap expiremap.ExpireMap[*tcpip.Address]
var writeVnetFun func([]byte)

var _ = divert.MustLoad(divert.DLL)

func SetWriteVNetFun(writeVnet func([]byte)) {
	writeVnetFun = writeVnet
}

func StartForward() error {
	ctx, cancel = context.WithCancel(context.Background())
	restart = make(chan bool)
	nats = []net.IPNet{}
	natMap.New(NAT_CACHE_TIMEOUT, 2*NAT_CACHE_TIMEOUT)

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return err
	}

	loaclIPMap = map[uint32]bool{}
	for _, addr := range addrs {
		netAddr := addr.(*net.IPNet)
		ipv4 := netAddr.IP.To4()
		if ipv4 == nil {
			continue
		}

		loaclIPMap[binary.BigEndian.Uint32(ipv4)] = true
	}

	defaultIP := getDefaultInterface()
	if len(defaultIP) == 0 {
		return errors.New("default interface not found")
	}

	defaultAddr = net.ParseIP(defaultIP).To4()

	wg.Add(1)
	go forward()

	return nil
}

func StopForward() {
	if ctx == nil {
		return
	}

	select {
	case <-ctx.Done():
		return
	default:
		cancel()
		wg.Wait()
		close(restart)
	}
}

func DelNatMasquerade(netAddr net.IPNet) {
	netAddr.IP = netAddr.IP.Mask(netAddr.Mask)
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
	netAddr.IP = netAddr.IP.Mask(netAddr.Mask)
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
	forward := Windivert{}
	network := Windivert{}

	closeDev := func() {
		network.Stop()
		forward.Stop()
	}

	restartDev := func() error {
		closeDev()

		needForward := func(iphdr *header.IPv4) bool {
			_, ok := loaclIPMap[binary.BigEndian.Uint32(iphdr.SourceAddressSlice())]
			if ok {
				return false
			}

			natsLock.RLock()
			defer natsLock.RUnlock()
			for _, nat := range nats {
				if nat.Contains(iphdr.SourceAddressSlice()) {
					return true
				}
			}

			return false
		}

		procForward := func(handle *divert.Handle, addr *divert.Address, pkg []byte) {
			if addr.IPv6() {
				handle.Send(pkg, addr)
				return
			}

			iphdr := header.IPv4(pkg)
			if !needForward(&iphdr) {
				handle.Send(pkg, addr)
				return
			}

			saddr := iphdr.SourceAddress()

			iphdr.SetSourceAddress(tcpip.AddrFrom4Slice(defaultAddr))
			calcChecksum(&iphdr, pkg)

			key := genSNatKey(&iphdr)
			oldAddr, expire := natMap.GetWithExpiration(key)
			if oldAddr == nil || NAT_CACHE_TIMEOUT-time.Since(expire).Abs() >= 1*time.Minute {
				natMap.Set(key, &saddr)
			}

			handle.Send(pkg, addr)
		}

		err := forward.Start("true", divert.NetworkForward, procForward)
		if err != nil {
			return err
		}

		procNetwork := func(handle *divert.Handle, addr *divert.Address, pkg []byte) {
			if addr.IPv6() {
				handle.Send(pkg, addr)
				return
			}

			if addr.Outbound() {
				handle.Send(pkg, addr)
				return
			}

			iphdr := header.IPv4(pkg)
			key := genDNatKey(&iphdr)
			oldAddr, expire := natMap.GetWithExpiration(key)
			if oldAddr == nil {
				handle.Send(pkg, addr)
				return
			}

			if NAT_CACHE_TIMEOUT-time.Since(expire).Abs() >= 1*time.Minute {
				natMap.Set(key, oldAddr)
			}

			iphdr.SetDestinationAddress(*oldAddr)

			calcChecksum(&iphdr, pkg)

			if writeVnetFun != nil {
				writeVnetFun(pkg)
			} else {
				handle.Send(pkg, addr)
			}
		}

		err = network.Start("!loopback", divert.Network, procNetwork)
		if err != nil {
			return err
		}

		return nil
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

			err := restartDev()
			if err != nil {
				time.Sleep(5 * time.Second)
				restart <- true
			}
		case <-ctx.Done():
			return
		}
	}
}

func genSNatKey(iphdr *header.IPv4) string {
	var sPort uint16 = 0
	var dPort uint16 = 0

	if iphdr.Protocol() == uint8(header.TCPProtocolNumber) {
		tcp := header.TCP(iphdr.Payload())
		sPort = tcp.SourcePort()
		dPort = tcp.DestinationPort()
	} else if iphdr.Protocol() == uint8(header.UDPProtocolNumber) {
		udp := header.UDP(iphdr.Payload())
		sPort = udp.SourcePort()
		dPort = udp.DestinationPort()
	}

	return calcKey(iphdr.SourceAddressSlice(), iphdr.DestinationAddressSlice(), sPort, dPort)
}

func genDNatKey(iphdr *header.IPv4) string {
	var sPort uint16 = 0
	var dPort uint16 = 0

	if iphdr.Protocol() == uint8(header.TCPProtocolNumber) {
		tcp := header.TCP(iphdr.Payload())
		sPort = tcp.SourcePort()
		dPort = tcp.DestinationPort()
	} else if iphdr.Protocol() == uint8(header.UDPProtocolNumber) {
		udp := header.UDP(iphdr.Payload())
		sPort = udp.SourcePort()
		dPort = udp.DestinationPort()
	}

	return calcKey(iphdr.DestinationAddressSlice(), iphdr.SourceAddressSlice(), dPort, sPort)
}

func calcKey(src []byte, dst []byte, sPort uint16, dPort uint16) string {
	key := make([]byte, 12)
	index := 0
	copy(key[index:], src)
	index += 4

	copy(key[index:], dst)
	index += 4

	binary.BigEndian.PutUint16(key[index:], sPort)
	index += 2

	binary.BigEndian.PutUint16(key[index:], dPort)
	index += 2

	return string(key)
}

func calcChecksum(iphdr *header.IPv4, pkg []byte) {
	ipPayload := iphdr.Payload()

	iphdr.SetChecksum(0)
	iphdr.SetChecksum(^checksum.Checksum(pkg[:iphdr.HeaderLength()], 0))

	if iphdr.Protocol() == uint8(header.TCPProtocolNumber) {
		tcp := header.TCP(ipPayload)
		tcp.SetChecksum(0)
		tcp.SetChecksum(^checksum.Combine(header.PseudoHeaderChecksum(
			header.TCPProtocolNumber,
			iphdr.SourceAddress(), iphdr.DestinationAddress(),
			uint16(len(ipPayload)),
		), checksum.Checksum(ipPayload, 0)))
	} else if iphdr.Protocol() == uint8(header.UDPProtocolNumber) {
		udp := header.UDP(ipPayload)
		udp.SetChecksum(0)
		udp.SetChecksum(^checksum.Combine(header.PseudoHeaderChecksum(
			header.UDPProtocolNumber,
			iphdr.SourceAddress(), iphdr.DestinationAddress(),
			uint16(len(ipPayload)),
		), checksum.Checksum(ipPayload, 0)))
	}
}

func getDefaultInterface() string {
	cmd := exec.Command("cmd", "/c", "route print 0.0.0.0")
	output, _ := cmd.CombinedOutput()
	re := regexp.MustCompile(`0.0.0.0\s+0.0.0.0\s+\d+\.\d+\.\d+\.\d+\s+(\S+)`)
	matches := re.FindStringSubmatch(string(output))
	if len(matches) > 1 {
		return matches[1]
	}

	return ""
}

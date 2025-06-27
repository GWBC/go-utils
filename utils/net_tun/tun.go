package nettun

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"

	"github.com/GWBC/go-utils/src/net_tun/netset"
	"github.com/GWBC/go-utils/src/pool"
	"golang.zx2c4.com/wireguard/tun"
)

type Tun struct {
	dev  tun.Device
	nset netset.Netset

	wg     sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc

	writeChanLock sync.RWMutex
	writeChan     chan *pool.Block

	dataPool *pool.BlockPool
}

func (t *Tun) Init(name string, payloadOffset int) error {
	t.ctx, t.cancel = context.WithCancel(context.Background())

	dev, err := tun.CreateTUN(name, MTU_SIZE)
	if err != nil {
		return err
	}

	err = t.nset.Init(name)
	if err != nil {
		return err
	}

	//设置MTU
	SetMTU(&t.nset, MTU_SIZE)

	t.dev = dev
	t.dataPool = pool.CreateBlockPool(MTU_SIZE+payloadOffset+512, payloadOffset)

	return nil
}

func (t *Tun) UnInit() {
	t.Stop()
}

func (t *Tun) Start(readFun TunDevReadFun, execptFun TunExceptionFun) {
	select {
	case <-t.ctx.Done():
		t.ctx, t.cancel = context.WithCancel(context.Background())
	default:
	}

	t.writePacket(execptFun)
	t.readPacket(readFun, execptFun)
}

func (t *Tun) Stop() {
	t.stop()
	t.Wait()
}

func (t *Tun) Wait() {
	t.wg.Wait()
}

func (t *Tun) SetAddrV4(addr string, mask string) error {
	var b1, b2, b3, b4 byte
	_, err := fmt.Sscanf(mask, "%d.%d.%d.%d", &b1, &b2, &b3, &b4)
	if err != nil {
		return err
	}

	ipnet := net.IPNet{
		IP:   net.ParseIP(addr),
		Mask: net.IPv4Mask(b1, b2, b3, b4),
	}

	err = t.nset.SetIPAddresses(ipnet)
	if err != nil {
		return err
	}

	return nil
}

func (t *Tun) AddRoutes(routes []netset.RouteInfo) error {
	err := t.nset.AddRoutes(routes)
	if err != nil {
		return err
	}

	return nil
}

func (t *Tun) Write(data *pool.Block) error {
	t.writeChanLock.RLock()
	defer t.writeChanLock.RUnlock()
	if t.writeChan == nil {
		return errors.New("tun is close")
	}

	t.writeChan <- data

	return nil
}

func (t *Tun) GetBlock() *pool.Block {
	return t.dataPool.Get()
}

func (t *Tun) stop() {
	select {
	case <-t.ctx.Done():
		return
	default:
		t.cancel()

		if t.dev != nil {
			t.dev.Close()
		}
	}
}

func (t *Tun) writePacket(execptFun TunExceptionFun) {
	t.writeChan = make(chan *pool.Block, 360)

	t.wg.Add(1)

	go func() {
		defer func() {
			t.stop()
			t.writeChanLock.Lock()
			close(t.writeChan)
			t.writeChan = nil
			t.writeChanLock.Unlock()
			t.wg.Done()
		}()

		//写入tunn
		var writeTunnFun = func(data *pool.Block) bool {
			defer data.Release()

			_, err := t.dev.Write([][]byte{data.Pkg}, data.PayloadOffset)
			if err != nil {
				select {
				case <-t.ctx.Done():
					break
				default:
					go execptFun(err)
				}

				return false
			}

			return true
		}

		for {
			select {
			case <-t.ctx.Done():
				return
			case data := <-t.writeChan:
				if !writeTunnFun(data) {
					return
				}
			}
		}
	}()
}

func (t *Tun) readPacket(readFun TunDevReadFun, execptFun TunExceptionFun) {
	t.wg.Add(1)

	go func() {
		defer func() {
			t.stop()
			t.wg.Done()
		}()

		var (
			bufCount = t.dev.BatchSize()
			sizes    = make([]int, bufCount)
			bufs     = make([][]byte, bufCount)
			datas    = make([]*pool.Block, bufCount)
		)

		readTunnFun := func() bool {
			defer func() {
				for i := range bufs {
					datas[i].Release()
				}
			}()

			//从池子中获取数据块
			for i := range bufs {
				datas[i] = t.dataPool.Get()
				bufs[i] = datas[i].Pkg
			}

			count, err := t.dev.Read(bufs, sizes, t.dataPool.PayloadOffset)
			if err != nil {
				select {
				case <-t.ctx.Done():
					break
				default:
					go execptFun(err)
				}

				return false
			}

			for i := range count {
				readFun(datas[i].SetPkgSize(sizes[i] + t.dataPool.PayloadOffset).AddRef())
			}

			return true
		}

		for {
			if !readTunnFun() {
				break
			}
		}
	}()
}

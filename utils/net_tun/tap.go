package nettun

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"

	"github.com/GWBC/go-utils/src/net_tun/netset"
	"github.com/GWBC/go-utils/src/pool"
	"github.com/songgao/water"
)

type Tap struct {
	dev  *water.Interface
	nset netset.Netset

	wg     sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc

	writeChanLock sync.RWMutex
	writeChan     chan *pool.Block

	dataPool *pool.BlockPool
}

// Window上需要TAP驱动：https://github.com/OpenVPN/tap-windows6
// 安装：devcon.exe install tap0901.inf tap0901
// 删除：devcon.exe remove tap0901
// 缺陷：不能设置过大的MTU
func (t *Tap) Init(_ string, payloadOffset int) error {
	t.ctx, t.cancel = context.WithCancel(context.Background())

	dev, err := water.New(water.Config{
		DeviceType: water.TAP,
	})
	if err != nil {
		return err
	}

	err = t.nset.Init(dev.Name())
	if err != nil {
		return err
	}

	err = SetMTU(&t.nset, MTU_SIZE)
	if err != nil {
		MTU_SIZE >>= 1
		err := SetMTU(&t.nset, MTU_SIZE)
		if err != nil {
			MTU_SIZE = 1500
			SetMTU(&t.nset, MTU_SIZE)
		}
	}

	t.dataPool = pool.CreateBlockPool(MTU_SIZE+payloadOffset+512, payloadOffset)
	t.dev = dev

	return nil
}

func (t *Tap) UnInit() {
	t.Stop()
}

func (t *Tap) Start(readFun TunDevReadFun, execptFun TunExceptionFun) {
	select {
	case <-t.ctx.Done():
		t.ctx, t.cancel = context.WithCancel(context.Background())
	default:
	}

	t.writePacket(execptFun)
	t.readPacket(readFun, execptFun)
}

func (t *Tap) Stop() {
	t.stop()
	t.Wait()
}

func (t *Tap) Wait() {
	t.wg.Wait()
}

func (t *Tap) SetAddrV4(addr string, mask string) error {
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

func (t *Tap) AddRoutes(routes []netset.RouteInfo) error {
	err := t.nset.AddRoutes(routes)
	if err != nil {
		return err
	}

	return nil
}

func (t *Tap) Write(data *pool.Block) error {
	t.writeChanLock.RLock()
	defer t.writeChanLock.RUnlock()
	if t.writeChan == nil {
		return errors.New("tap is close")
	}

	t.writeChan <- data

	return nil
}

func (t *Tap) GetBlock() *pool.Block {
	return t.dataPool.Get()
}

func (t *Tap) stop() {
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

func (t *Tap) writePacket(execptFun TunExceptionFun) {
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

			_, err := t.dev.Write(data.Payload())
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

func (t *Tap) readPacket(readFun TunDevReadFun, execptFun TunExceptionFun) {
	t.wg.Add(1)

	go func() {
		defer func() {
			t.stop()
			t.wg.Done()
		}()

		readTunnFun := func() bool {
			data := t.dataPool.Get()
			defer data.Release()

			n, err := t.dev.Read(data.Payload())
			if err != nil {
				select {
				case <-t.ctx.Done():
					break
				default:
					go execptFun(err)
				}

				return false
			}

			readFun(data.SetPkgSize(n + t.dataPool.PayloadOffset).AddRef())

			return true
		}

		for {
			if !readTunnFun() {
				break
			}
		}
	}()
}

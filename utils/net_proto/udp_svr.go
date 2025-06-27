package netproto

import (
	"context"
	"errors"
	"net"

	"github.com/GWBC/go-utils/src/pool"
	"gvisor.dev/gvisor/pkg/sync"
)

type UDPConn struct {
	Svr      *UDPSvr
	ReadChan chan *pool.Block
	Addr     *net.UDPAddr
	NetRead  NetworkRead
	NetWrite NetworkWrite

	Ctx       context.Context
	CancelFun context.CancelFunc
}

func (u *UDPConn) Write(data *pool.Block) error {
	return u.NetWrite.Write(NetData{u.Addr, data})
}

func (u *UDPConn) Close() {
	select {
	case <-u.Ctx.Done():
		return
	default:
		u.CancelFun()
		close(u.ReadChan)
		u.Svr.delete(u.Addr.String())
	}
}

func (u *UDPConn) String() string {
	return u.Addr.String()
}

////////////////////////////////////////////////////////////////

type UDPSvr struct {
	SystemContext
	ServerCallback

	sock *net.UDPConn

	lock  sync.Mutex
	conns map[string]*UDPConn
}

func (u *UDPSvr) Start() error {
	u.conns = map[string]*UDPConn{}

	addr, err := net.ResolveUDPAddr("udp4", u.addr)
	if err != nil {
		return err
	}

	sock, err := net.ListenUDP("udp4", addr)
	if err != nil {
		return err
	}

	u.wg.Add(1)

	go func() {
		defer func() {
			u.close()
			u.wg.Done()
		}()

		for {
			block := u.dataPool.Get()

			//读取网络包
			n, addr, err := sock.ReadFromUDP(block.Pkg)
			if err != nil {
				return
			}

			block.SetPkgSize(n)
			strAddr := addr.String()

			u.lock.Lock()
			conn := u.conns[strAddr]

			if conn == nil {
				conn = u.newConn(addr)
				u.conns[strAddr] = conn
			}

			conn.ReadChan <- block
			u.lock.Unlock()
		}
	}()

	u.sock = sock

	return nil
}

func (u *UDPSvr) Stop() {
	u.close()
	u.Wait()
}

func (u *UDPSvr) Wait() {
	u.wg.Wait()
}

func (u *UDPSvr) newConn(addr *net.UDPAddr) *UDPConn {
	connObj := &UDPConn{Svr: u, Addr: addr}
	connObj.Ctx, connObj.CancelFun = context.WithCancel(u.ctx)

	connObj.ReadChan = make(chan *pool.Block, 360)

	stop := func() {
		u.delete(addr.String())
	}

	except := func(addr net.Addr, err error) {
		u.exceptFun(connObj, err)
	}

	connObj.NetRead.SetContext(u.ctx)
	connObj.NetRead.Start(ReadStartInfo{
		Conn:        connObj,
		Group:       &u.wg,
		DataPool:    u.dataPool,
		ReadPayload: u.rPayload,
		HeartCheck:  u.newHeartCheck(connObj),
		Read: func(data []byte) (int, net.Addr, error) {
			select {
			case block := <-connObj.ReadChan:
				defer block.Release()

				if block == nil {
					return 0, addr, errors.New("read is close")
				}

				copy(data, block.Pkg)
				return len(block.Pkg), addr, nil
			case <-connObj.Ctx.Done():
				return 0, addr, errors.New("read is eof")
			}
		},
		Except:       except,
		StopCallback: stop,
		ReadCallback: u.readFun,
	})

	connObj.NetWrite.SetChanSize(360).SetContext(u.ctx)
	connObj.NetWrite.Start(WriteStartInfo{
		Group:        &u.wg,
		WritePayload: u.wPayload,
		Write: func(addr net.Addr, data []byte) (int, error) {
			return u.sock.WriteToUDP(data, addr.(*net.UDPAddr))
		},
		Except:       except,
		StopCallback: stop,
	})

	return connObj
}

func (u *UDPSvr) delete(addr string) {
	u.lock.Lock()
	defer u.lock.Unlock()
	delete(u.conns, addr)
}

func (u *UDPSvr) close() {
	select {
	case <-u.ctx.Done():
		return
	default:
		if u.sock == nil {
			return
		}

		u.sock.Close()
		u.cancel()
	}
}

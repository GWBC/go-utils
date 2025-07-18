package netproto

import (
	"context"
	"errors"
	"net"

	expiremap "github.com/GWBC/go-utils/utils/expire_map"
	"github.com/GWBC/go-utils/utils/pool"
)

type UDPConn struct {
	NetworkUserData
	Svr      *UDPSvr
	ReadChan chan *pool.Block
	LAddr    *net.UDPAddr
	RAddr    *net.UDPAddr
	NetRead  NetworkRead
	NetWrite NetworkWrite

	Ctx       context.Context
	CancelFun context.CancelFunc
}

func (u *UDPConn) Write(data *pool.Block) error {
	return u.NetWrite.Write(NetData{u.RAddr, data})
}

func (u *UDPConn) Close() {
	select {
	case <-u.Ctx.Done():
		return
	default:
		u.CancelFun()
		u.Svr.delete(u.RAddr.String())
	}
}

func (u *UDPConn) LocalAddr() string {
	return u.LAddr.String()
}

func (u *UDPConn) RemoteAddr() string {
	return u.RAddr.String()
}

////////////////////////////////////////////////////////////////

type UDPSvr struct {
	SystemContext
	ServerCallback
	sock  *net.UDPConn
	conns expiremap.ExpireMap[*UDPConn]
}

func (u *UDPSvr) Start() error {
	if len(u.netType) == 0 {
		u.netType = "UDP"
	}

	u.conns.New(-1, -1)

	laddr, err := net.ResolveUDPAddr("udp4", u.addr)
	if err != nil {
		return err
	}

	sock, err := net.ListenUDP("udp4", laddr)
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
			n, raddr, err := sock.ReadFromUDP(block.Pkg)
			if err != nil {
				return
			}

			block.SetPkgSize(n)
			remoteAddr := raddr.String()

			conn := u.conns.Get(remoteAddr)
			if conn == nil {
				conn = u.newConn(raddr, laddr)
				u.conns.Set(remoteAddr, conn)
			}

			conn.ReadChan <- block
		}
	}()

	u.sock = sock

	return nil
}

func (u *UDPSvr) Stop() {
	if u.ctx == nil {
		return
	}

	u.close()
	u.Wait()
}

func (u *UDPSvr) Wait() {
	u.wg.Wait()
}

func (u *UDPSvr) newConn(raddr *net.UDPAddr, laddr *net.UDPAddr) *UDPConn {
	connObj := &UDPConn{Svr: u, RAddr: raddr, LAddr: laddr}
	connObj.Ctx, connObj.CancelFun = context.WithCancel(u.ctx)
	connObj.ReadChan = make(chan *pool.Block, 360)

	stop := func() {
		u.delete(raddr.String())
	}

	except := func(addr net.Addr, err error) {
		u.exceptFun(connObj, err)
	}

	connObj.NetWrite.SetChanSize(360).SetContext(u.ctx)
	connObj.NetRead.SetContext(u.ctx)

	connObj.NetWrite.Start(WriteStartInfo{
		Group:        &u.wg,
		WritePayload: u.wPayload,
		Write: func(addr net.Addr, data []byte) (int, error) {
			return u.sock.WriteToUDP(data, addr.(*net.UDPAddr))
		},
		Except:       except,
		StopCallback: stop,
	})

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
					return 0, raddr, errors.New("read is close")
				}

				copy(data, block.Pkg)
				return len(block.Pkg), raddr, nil
			case <-connObj.Ctx.Done():
				return 0, raddr, errors.New("read is eof")
			}
		},
		Except:       except,
		StopCallback: stop,
		ReadCallback: u.readFun,
	})

	return connObj
}

func (u *UDPSvr) delete(addr string) {
	u.conns.Delete(addr)
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

package netproto

import (
	"net"

	"github.com/GWBC/go-utils/utils/pool"
)

type UDPClient struct {
	NetworkUserData
	SystemContext
	ClientCallback

	sock     *net.UDPConn
	netRead  NetworkRead
	netWrite NetworkWrite
}

func (u *UDPClient) Start() error {
	addr, err := net.ResolveUDPAddr("udp4", u.addr)
	if err != nil {
		return err
	}

	sock, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		return err
	}

	u.sock = sock

	except := func(addr net.Addr, err error) {
		u.exceptFun(u, err)
	}

	u.netWrite.SetChanSize(360).SetContext(u.ctx)
	u.netRead.SetContext(u.ctx).SetDecode(u.decodes)

	u.netWrite.Start(WriteStartInfo{
		Group:        &u.wg,
		WritePayload: u.wPayload,
		Write: func(addr net.Addr, data []byte) (int, error) {
			return u.sock.Write(data)
		},
		Except:       except,
		StopCallback: u.Close,
	})

	u.netRead.Start(ReadStartInfo{
		Conn:        u,
		Group:       &u.wg,
		DataPool:    u.dataPool,
		ReadPayload: u.rPayload,
		HeartCheck:  u.newHeartCheck(u),
		Read: func(data []byte) (int, net.Addr, error) {
			return u.sock.ReadFromUDP(data)
		},
		Except:       except,
		StopCallback: u.Close,
		ReadCallback: u.readFun,
	})

	return nil
}

func (u *UDPClient) Stop() {
	if u.ctx == nil {
		return
	}

	u.Close()
	u.Wait()
}

func (u *UDPClient) Wait() {
	u.wg.Wait()
}

func (u *UDPClient) Write(data *pool.Block) error {
	return u.netWrite.Write(NetData{nil, data})
}

func (u *UDPClient) Close() {
	select {
	case <-u.ctx.Done():
		return
	default:
		if u.sock == nil {
			return
		}

		u.cancel()
		u.sock.Close()
	}
}

func (u *UDPClient) LocalAddr() string {
	return u.sock.LocalAddr().String()
}

func (u *UDPClient) RemoteAddr() string {
	return u.sock.RemoteAddr().String()
}

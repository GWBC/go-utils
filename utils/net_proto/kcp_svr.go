package netproto

import (
	"net"
	"sync"

	"github.com/GWBC/go-utils/utils/pool"
	"github.com/xtaci/kcp-go/v5"
)

type KCPConn struct {
	NetworkUserData
	Conn     net.Conn
	NetRead  NetworkRead
	NetWrite NetworkWrite
}

func (k *KCPConn) Write(data *pool.Block) error {
	return k.NetWrite.Write(NetData{nil, data})
}

func (k *KCPConn) Close() {
	k.Conn.Close()
}

func (k *KCPConn) LocalAddr() string {
	return k.Conn.LocalAddr().String()
}

func (k *KCPConn) RemoteAddr() string {
	return k.Conn.RemoteAddr().String()
}

////////////////////////////////////////////////////////////////

type KCPSvr struct {
	SystemContext
	ServerCallback

	sock *kcp.Listener

	lock  sync.RWMutex
	conns map[*KCPConn]*KCPConn
}

func (k *KCPSvr) Start() error {
	//FEC前向纠错：10包允许丢3个包，丢包率30%，带宽相应增加30%
	sock, err := kcp.ListenWithOptions(k.addr, nil, 10, 3)
	if err != nil {
		return err
	}

	k.sock = sock
	k.conns = map[*KCPConn]*KCPConn{}

	k.accept()

	return nil
}

func (k *KCPSvr) Stop() {
	k.stop()
	k.Wait()
}

func (k *KCPSvr) stop() {
	if k.ctx == nil {
		return
	}

	select {
	case <-k.ctx.Done():
		return
	default:
		if k.sock == nil {
			return
		}

		k.cancel()
		k.sock.Close()

		k.lock.RLock()
		defer k.lock.RUnlock()
		for _, conn := range k.conns {
			conn.Close()
		}
	}
}

func (k *KCPSvr) accept() {
	k.wg.Add(1)

	go func() {
		defer func() {
			k.stop()
			k.wg.Done()
		}()

		for {
			conn, err := k.sock.AcceptKCP()
			if err != nil {
				select {
				case <-k.ctx.Done():
					break
				default:
					go k.exceptFun(nil, err)
				}

				break
			}

			k.newConn(conn)
		}
	}()
}

func (k *KCPSvr) newConn(conn *kcp.UDPSession) {
	connObj := &KCPConn{Conn: conn}

	k.lock.Lock()
	k.conns[connObj] = connObj
	k.lock.Unlock()

	stop := func() {
		k.lock.Lock()
		defer k.lock.Unlock()
		connObj.Close()
		delete(k.conns, connObj)
	}

	except := func(addr net.Addr, err error) {
		k.exceptFun(connObj, err)
	}

	connObj.NetWrite.SetChanSize(360).SetContext(k.ctx)
	connObj.NetRead.SetContext(k.ctx).SetDecode(k.decodes)

	connObj.NetWrite.Start(WriteStartInfo{
		Group:        &k.wg,
		WritePayload: k.wPayload,
		Write: func(addr net.Addr, data []byte) (int, error) {
			return conn.Write(data)
		},
		Except:       except,
		StopCallback: stop,
	})

	connObj.NetRead.Start(ReadStartInfo{
		Conn:        connObj,
		Group:       &k.wg,
		DataPool:    k.dataPool,
		ReadPayload: k.rPayload,
		HeartCheck:  k.newHeartCheck(connObj),
		Read: func(data []byte) (int, net.Addr, error) {
			n, err := conn.Read(data)
			return n, nil, err
		},
		Except:       except,
		StopCallback: stop,
		ReadCallback: k.readFun,
	})

}

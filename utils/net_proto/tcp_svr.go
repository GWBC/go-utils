package netproto

import (
	"net"
	"sync"

	"github.com/GWBC/go-utils/utils/pool"
)

type TCPConn struct {
	Conn     net.Conn
	NetRead  NetworkRead
	NetWrite NetworkWrite
}

func (t *TCPConn) Write(data *pool.Block) error {
	return t.NetWrite.Write(NetData{nil, data})
}

func (t *TCPConn) Close() {
	t.Conn.Close()
}

func (t *TCPConn) String() string {
	return t.Conn.RemoteAddr().String()
}

////////////////////////////////////////////////////////////////

type TCPSvr struct {
	SystemContext
	ServerCallback

	sock net.Listener

	lock  sync.RWMutex
	conns map[*TCPConn]*TCPConn
}

func (t *TCPSvr) Start() error {
	sock, err := net.Listen("tcp", t.addr)
	if err != nil {
		return err
	}

	t.sock = sock
	t.conns = map[*TCPConn]*TCPConn{}

	t.accept()

	return nil
}

func (t *TCPSvr) Stop() {
	if t.ctx == nil {
		return
	}

	t.stop()
	t.Wait()
}

func (t *TCPSvr) stop() {
	select {
	case <-t.ctx.Done():
		return
	default:
		if t.sock == nil {
			return
		}

		t.cancel()
		t.sock.Close()

		t.lock.RLock()
		defer t.lock.RUnlock()
		for _, conn := range t.conns {
			conn.Close()
		}
	}
}

func (t *TCPSvr) accept() {
	t.wg.Add(1)

	go func() {
		defer func() {
			t.stop()
			t.wg.Done()
		}()

		for {
			conn, err := t.sock.Accept()
			if err != nil {
				select {
				case <-t.ctx.Done():
					break
				default:
					go t.exceptFun(nil, err)
				}

				break
			}

			t.newConn(conn)
		}
	}()
}

func (t *TCPSvr) newConn(conn net.Conn) {
	connObj := &TCPConn{Conn: conn}

	t.lock.Lock()
	t.conns[connObj] = connObj
	t.lock.Unlock()

	stop := func() {
		t.lock.Lock()
		defer t.lock.Unlock()
		connObj.Close()
		delete(t.conns, connObj)
	}

	except := func(addr net.Addr, err error) {
		t.exceptFun(connObj, err)
	}

	connObj.NetRead.SetContext(t.ctx).SetDecode(t.decodes)
	connObj.NetRead.Start(ReadStartInfo{
		Conn:        connObj,
		Group:       &t.wg,
		DataPool:    t.dataPool,
		ReadPayload: t.rPayload,
		HeartCheck:  t.newHeartCheck(connObj),
		Read: func(data []byte) (int, net.Addr, error) {
			n, err := conn.Read(data)
			return n, nil, err
		},
		Except:       except,
		StopCallback: stop,
		ReadCallback: t.readFun,
	})

	connObj.NetWrite.SetChanSize(360).SetContext(t.ctx)
	connObj.NetWrite.Start(WriteStartInfo{
		Group:        &t.wg,
		WritePayload: t.wPayload,
		Write: func(addr net.Addr, data []byte) (int, error) {
			return conn.Write(data)
		},
		Except:       except,
		StopCallback: stop,
	})
}

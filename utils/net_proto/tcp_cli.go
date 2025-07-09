package netproto

import (
	"net"
	"time"

	"github.com/GWBC/go-utils/utils/pool"
)

type TCPClient struct {
	SystemContext
	ClientCallback

	conn     net.Conn
	netRead  NetworkRead
	netWrite NetworkWrite
}

func (t *TCPClient) Start() error {
	conn, err := net.DialTimeout("tcp", t.addr, 10*time.Second)
	if err != nil {
		return err
	}

	t.conn = conn

	except := func(addr net.Addr, err error) {
		t.exceptFun(t, err)
	}

	t.netRead.SetContext(t.ctx).SetDecode(t.decodes)
	t.netRead.Start(ReadStartInfo{
		Conn:        t,
		Group:       &t.wg,
		DataPool:    t.dataPool,
		ReadPayload: t.rPayload,
		HeartCheck:  t.newHeartCheck(t),
		Read: func(data []byte) (int, net.Addr, error) {
			n, err := t.conn.Read(data)
			return n, nil, err
		},
		Except:       except,
		StopCallback: t.Close,
		ReadCallback: t.readFun,
	})

	t.netWrite.SetChanSize(360).SetContext(t.ctx)
	t.netWrite.Start(WriteStartInfo{
		Group:        &t.wg,
		WritePayload: t.wPayload,
		Write: func(addr net.Addr, data []byte) (int, error) {
			return conn.Write(data)
		},
		Except:       except,
		StopCallback: t.Close,
	})

	return nil
}

func (t *TCPClient) Stop() {
	if t.ctx == nil {
		return
	}

	t.Close()
	t.Wait()
}

func (t *TCPClient) Write(data *pool.Block) error {
	return t.netWrite.Write(NetData{Addr: nil, Data: data})
}

func (t *TCPClient) Close() {
	select {
	case <-t.ctx.Done():
		return
	default:
		if t.conn == nil {
			return
		}

		t.cancel()
		t.conn.Close()
	}
}

func (t *TCPClient) String() string {
	return t.conn.RemoteAddr().String()
}

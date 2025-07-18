package netproto

import (
	"net"

	"github.com/GWBC/go-utils/utils/pool"
	"github.com/xtaci/kcp-go/v5"
)

type KCPClient struct {
	NetworkUserData
	SystemContext
	ClientCallback

	conn     net.Conn
	netRead  NetworkRead
	netWrite NetworkWrite
}

func (k *KCPClient) Start() error {
	if len(k.netType) == 0 {
		k.netType = "KCP"
	}

	k.SetData(nil)

	//FEC前向纠错：10包允许丢3个包，丢包率30%，带宽相应增加30%
	conn, err := kcp.DialWithOptions(k.addr, nil, 10, 3)
	if err != nil {
		return err
	}

	k.conn = conn

	except := func(addr net.Addr, err error) {
		k.exceptFun(k, err)
	}

	k.netWrite.SetChanSize(360).SetContext(k.ctx)
	k.netRead.SetContext(k.ctx).SetDecode(k.decodes)

	k.netWrite.Start(WriteStartInfo{
		Group:        &k.wg,
		WritePayload: k.wPayload,
		Write: func(addr net.Addr, data []byte) (int, error) {
			return conn.Write(data)
		},
		Except:       except,
		StopCallback: k.Close,
		Hook:         k.wHook,
	})

	k.netRead.Start(ReadStartInfo{
		Conn:        k,
		Group:       &k.wg,
		DataPool:    k.dataPool,
		ReadPayload: k.rPayload,
		HeartCheck:  k.newHeartCheck(k),
		Read: func(data []byte) (int, net.Addr, error) {
			n, err := k.conn.Read(data)
			return n, nil, err
		},
		Except:       except,
		StopCallback: k.Close,
		ReadCallback: k.readFun,
		Hook:         k.rHook,
	})

	return nil
}

func (k *KCPClient) Stop() {
	if k.ctx == nil {
		return
	}

	k.Close()
	k.Wait()
}

func (k *KCPClient) Write(data *pool.Block) error {
	return k.netWrite.Write(NetData{Addr: nil, Data: data})
}

func (k *KCPClient) Close() {
	if k.ctx == nil {
		return
	}

	select {
	case <-k.ctx.Done():
		return
	default:
		if k.conn == nil {
			return
		}

		k.cancel()
		k.conn.Close()
	}
}

func (k *KCPClient) LocalAddr() string {
	return k.conn.LocalAddr().String()
}

func (k *KCPClient) RemoteAddr() string {
	return k.conn.RemoteAddr().String()
}

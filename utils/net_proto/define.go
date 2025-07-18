package netproto

import (
	"context"
	"net"
	"time"

	"github.com/GWBC/go-utils/utils/pool"
)

type HeartFun = func(conn Connection)
type NewHeartFun = func(conn Connection) HeartbeatCheck
type ExceptionFun = func(conn Connection, err error)
type NetReadFun = func(conn Connection, addr net.Addr, data *pool.Block)

type HeartbeatCheckInfo struct {
	Ctx               context.Context //上下文
	Conn              Connection      //连接对象
	SendTime          time.Duration   //发送心跳间隔
	CheckTime         time.Duration   //检查心跳间隔
	SendHeartCallback HeartFun        //发送心跳回调
	TimeoutCallback   HeartFun        //心跳超时回调
}

// 心跳检测
type HeartbeatCheck interface {
	Start(info HeartbeatCheckInfo) error
	Stop()
	Update()
}

// 流解析器
type StreamDecode interface {
	New() StreamDecode
	Decode(data *pool.Block) ([]*pool.Block, error)
}

// 用户数据接口
type UserData interface {
	SetData(data any)
	GetData() any
}

// 连接对象
type Connection interface {
	UserData
	Write(data *pool.Block) error
	Close()
	LocalAddr() string
	RemoteAddr() string
}

// 网络环境
type NetworkContext interface {
	NewContext() NetworkContext

	GetType() string

	SetAddr(addr string) NetworkContext
	GetAddr() string

	NewHeartCheck(fun NewHeartFun) NetworkContext
	AddDecode(decode StreamDecode) NetworkContext

	GetBlock() *pool.Block
	SetBlock(size int, payloadOffset int) NetworkContext
	SetRWInfo(readPayload bool, writePayload bool) NetworkContext

	Wait()
}

// 网络基础
type NetworkBase interface {
	NetworkContext
	SetReadCallback(readFun NetReadFun)
	SetExceptionCallback(exceptFun ExceptionFun)
	Start() error
	Stop()
}

// 服务端
type NetworkServer interface {
	NetworkBase
}

// 客户端
type NetworkClient interface {
	NetworkBase
	Connection
}

package netproto

import (
	"context"
	"errors"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/GWBC/go-utils/utils/pool"
)

/////////////////////////////////////////////////

type NetworkUserData struct {
	data any
}

func (n *NetworkUserData) SetData(data any) {
	n.data = data
}

func (n *NetworkUserData) GetData() any {
	return n.data
}

/////////////////////////////////////////////////

type NetworkHeartCheck struct {
	count atomic.Uint32

	ctx    context.Context
	cancel context.CancelFunc
}

func (n *NetworkHeartCheck) Start(info HeartbeatCheckInfo) error {
	n.ctx, n.cancel = context.WithCancel(info.Ctx)
	sendTicker := time.NewTicker(info.SendTime)
	checkTicker := time.NewTicker(info.CheckTime)

	go func() {
		defer func() {
			sendTicker.Stop()
			checkTicker.Stop()
		}()

		for {
			select {
			case <-sendTicker.C:
				if info.SendHeartCallback != nil {
					info.SendHeartCallback(info.Conn)
				}
			case <-checkTicker.C:
				if n.count.Swap(0) == 0 {
					info.TimeoutCallback(info.Conn)
				}
			case <-n.ctx.Done():
				return
			}
		}
	}()

	return nil
}

func (n *NetworkHeartCheck) Stop() {
	select {
	case <-n.ctx.Done():
		break
	default:
		n.cancel()
	}
}

func (n *NetworkHeartCheck) Update() {
	n.count.Add(1)
}

/////////////////////////////////////////////////

type SystemContext struct {
	wg     sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc

	dataPool *pool.BlockPool

	addr    string
	netType string

	decodes       []StreamDecode
	newHeartCheck NewHeartFun

	rHook HookFun
	wHook HookFun
}

func (s *SystemContext) NewContext() NetworkContext {
	s.ctx, s.cancel = context.WithCancel(context.Background())
	return s
}

func (s *SystemContext) HookRead(fun HookFun) NetworkContext {
	s.rHook = fun
	return s
}

func (s *SystemContext) HookWrite(fun HookFun) NetworkContext {
	s.wHook = fun
	return s
}

func (s *SystemContext) SetType(netType string) {
	s.netType = netType
}

func (s *SystemContext) GetType() string {
	return s.netType
}

func (s *SystemContext) SetAddr(addr string) NetworkContext {
	s.addr = addr
	return s
}

func (s *SystemContext) GetAddr() string {
	return s.addr
}

func (s *SystemContext) AddDecode(decode StreamDecode) NetworkContext {
	s.decodes = append(s.decodes, decode.New())
	return s
}

func (s *SystemContext) NewHeartCheck(fun NewHeartFun) NetworkContext {
	s.newHeartCheck = fun
	return s
}

func (s *SystemContext) GetBlock() *pool.Block {
	return s.dataPool.Get()
}

func (s *SystemContext) SetBlock(size int, payloadOffset int) NetworkContext {
	s.dataPool = pool.CreateBlockPool(size, payloadOffset)
	return s
}

func (s *SystemContext) Wait() {
	s.wg.Wait()
}

/////////////////////////////////////////////////

type ClientCallback struct {
	readFun   NetReadFun
	exceptFun ExceptionFun
}

func (s *ClientCallback) SetReadCallback(readFun NetReadFun) {
	s.readFun = readFun
}

func (s *ClientCallback) SetExceptionCallback(exceptFun ExceptionFun) {
	s.exceptFun = exceptFun
}

/////////////////////////////////////////////////

type ServerCallback struct {
	readFun   NetReadFun
	exceptFun ExceptionFun
}

func (s *ServerCallback) SetReadCallback(readFun NetReadFun) {
	s.readFun = readFun
}

func (s *ServerCallback) SetExceptionCallback(exceptFun ExceptionFun) {
	s.exceptFun = exceptFun
}

/////////////////////////////////////////////////

type NetData struct {
	Addr net.Addr
	Data *pool.Block
}

type WriteStartInfo struct {
	Group        *sync.WaitGroup
	Write        func(addr net.Addr, data []byte) (int, error)
	Except       func(addr net.Addr, err error)
	StopCallback func()
	Hook         HookFun
}

type NetworkWrite struct {
	ctx       context.Context
	writeChan chan NetData
}

func (n *NetworkWrite) SetContext(ctx context.Context) *NetworkWrite {
	n.ctx = ctx
	return n
}

func (n *NetworkWrite) SetChanSize(size int) *NetworkWrite {
	n.writeChan = make(chan NetData, size)
	return n
}

func (n *NetworkWrite) Write(data NetData) error {
	select {
	case <-n.ctx.Done():
		return errors.New("network write is close")
	default:
		n.writeChan <- data
	}

	return nil
}

func (n *NetworkWrite) Start(info WriteStartInfo) {
	info.Group.Add(1)

	go func() {
		defer func() {
			if info.StopCallback != nil {
				info.StopCallback()
			}

			info.Group.Done()
		}()

		dWrite := func(block *NetData) (int, error) {
			return info.Write(block.Addr, block.Data.Pkg)
		}

		if info.Hook != nil {
			dWrite = func(block *NetData) (int, error) {
				info.Hook(block.Data)
				return info.Write(block.Addr, block.Data.Pkg)
			}
		}

		write := func(block *NetData) bool {
			defer block.Data.Release()
			_, err := dWrite(block)
			if err != nil {
				select {
				case <-n.ctx.Done():
					break
				default:
					go info.Except(block.Addr, err)
				}

				return false
			}

			return true
		}

		for {
			select {
			case data := <-n.writeChan:
				if !write(&data) {
					return
				}
			case <-n.ctx.Done():
				return
			}
		}
	}()
}

/////////////////////////////////////////////////

type ReadStartInfo struct {
	Conn         Connection
	Group        *sync.WaitGroup
	DataPool     *pool.BlockPool
	HeartCheck   HeartbeatCheck
	Read         func(data []byte) (int, net.Addr, error)
	Except       func(addr net.Addr, err error)
	StopCallback func()
	ReadCallback NetReadFun
	Hook         HookFun
}

type NetworkRead struct {
	ctx     context.Context
	decodes []StreamDecode
}

func (n *NetworkRead) SetDecode(decodes []StreamDecode) *NetworkRead {
	for _, d := range decodes {
		n.decodes = append(n.decodes, d.New())
	}

	return n
}

func (n *NetworkRead) SetContext(ctx context.Context) *NetworkRead {
	n.ctx = ctx
	return n
}

func (n *NetworkRead) Start(info ReadStartInfo) {
	info.Group.Add(1)

	go func() {
		defer func() {
			if info.HeartCheck != nil {
				info.HeartCheck.Stop()
			}

			if info.StopCallback != nil {
				info.StopCallback()
			}

			info.Group.Done()
		}()

		tmpExceptFun := func(addr net.Addr, err error) {
			select {
			case <-n.ctx.Done():
				break
			default:
				go info.Except(addr, err)
			}
		}

		readCallBack := info.ReadCallback
		if info.Hook != nil {
			readCallBack = func(conn Connection, addr net.Addr, data *pool.Block) {
				info.Hook(data)
				info.ReadCallback(conn, addr, data)
			}
		}

		read := func() bool {
			block := info.DataPool.Get()
			defer block.Release()

			//读取网络包
			rn, addr, err := info.Read(block.Pkg)
			if err != nil {
				tmpExceptFun(addr, err)
				return false
			}

			block.SetPkgSize(rn)

			if info.HeartCheck != nil {
				info.HeartCheck.Update()
			}

			//解码对象接管
			if len(n.decodes) != 0 {
				for _, decode := range n.decodes {
					blocks, err := decode.Decode(block.AddRef())
					if err != nil {
						tmpExceptFun(addr, err)
						return false
					}

					for _, data := range blocks {
						readCallBack(info.Conn, addr, data.AddRef())
					}
				}

				return true
			}

			readCallBack(info.Conn, addr, block.AddRef())

			return true
		}

		for {
			if !read() {
				return
			}
		}
	}()
}

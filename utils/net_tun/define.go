package nettun

import (
	"time"

	"github.com/GWBC/go-utils/utils/net_tun/netset"
	"github.com/GWBC/go-utils/utils/pool"
)

var MTU_SIZE = 8 * 1024

const VirtioNetHdrLen = 10

type TunExceptionFun = func(err error)
type TunDevReadFun = func(data *pool.Block)

type VNetDev interface {
	Init(name string, payloadOffset int) error
	UnInit()
	Wait()

	Start(readFun TunDevReadFun, execptFun TunExceptionFun)
	Stop()

	SetAddrV4(addr string, mask string) error
	AddRoutes(routes []netset.RouteInfo) error

	GetBlock() *pool.Block
	Write(data *pool.Block) error
}

func SetMTU(nset *netset.Netset, size int) error {
	var err error

	for range 5 {
		err = nset.SetMTU(size)
		if err == nil {
			break
		}

		time.Sleep(200 * time.Millisecond)
	}

	return err
}

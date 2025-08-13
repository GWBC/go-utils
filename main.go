package main

import (
	"fmt"
	"net"
	"path/filepath"
	"sync"
	"time"

	"github.com/GWBC/go-utils/utils"
	jsengine "github.com/GWBC/go-utils/utils/js_engine"
	"github.com/GWBC/go-utils/utils/net_tun/netset"
)

func NatTest() {
	netaddr := net.IPNet{}
	netaddr.IP = net.ParseIP("10.0.0.23")
	netaddr.Mask = net.IPv4Mask(255, 255, 255, 0)
	netset.StartForward()
	netset.DelNatMasquerade(netaddr)
	netset.AddNatMasquerade(netaddr)

	time.Sleep(1000 * time.Second)
	netset.StopForward()
}

func TestJS() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		fpath := filepath.Join(utils.Pwd(), "..", "douban.js")

		a, _ := jsengine.New("douban")
		err := a.Require(fpath, "home")
		if err != nil {
			panic(err)
		}

		ret, err := a.RunString("home.Home()")
		if err != nil {
			panic(err)
		}

		ret, err = a.RunString("home.Data('movie', '热门', 1, 30)")
		if err != nil {
			panic(err)
		}

		fmt.Println(ret.(string))
	}()

	wg.Wait()
}

func main() {
	//NatTest()

	TestJS()
}

package main

import (
	"fmt"
	"net"
	"sync"
	"time"

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

	wg.Add(2)

	go func() {
		defer wg.Done()

		a, _ := jsengine.New("t1")
		err := a.Require("./a.js", "t")
		if err != nil {
			panic(err)
		}

		ret, err := a.RunString("t.getDuoban()")
		if err != nil {
			panic(err)
		}

		fmt.Println(ret.(string))
	}()

	go func() {
		defer wg.Done()

		a, _ := jsengine.New("t1")
		err := a.Require("./a.js", "t")
		if err != nil {
			panic(err)
		}

		ret, err := a.RunString("t.getDuoban()")
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

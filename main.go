package main

import (
	"net"
	"time"

	"github.com/GWBC/go-utils/utils/net_tun/netset"
)

func main() {
	_, netAddr, _ := net.ParseCIDR("10.0.0.0/24")
	netset.StartForward()
	netset.DelNatMasquerade(*netAddr)
	netset.AddNatMasquerade(*netAddr)

	time.Sleep(10000 * time.Second)
	netset.StopForward()
}

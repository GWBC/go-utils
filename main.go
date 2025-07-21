package main

import (
	"net"
	"time"

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

func main() {
	NatTest()
}

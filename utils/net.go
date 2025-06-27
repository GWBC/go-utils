package utils

import (
	"fmt"
	"net"
)

func GenCIDR(ip string, mask string) string {
	msk := net.IPMask(net.ParseIP(mask).To4())
	size, _ := msk.Size()
	return fmt.Sprintf("%s/%d", ip, size)
}

func GenIPNet(ip string, mask string) *net.IPNet {
	ret := net.IPNet{}
	ret.IP = net.ParseIP(ip).To4()
	ret.Mask = net.IPMask(net.ParseIP(mask).To4())
	return &ret
}

func CalculateCIDRBroadcast(cidr string) net.IP {
	_, n, _ := net.ParseCIDR(cidr)
	return CalculateBroadcast(n)
}

func CalculateBroadcast(ipNet *net.IPNet) net.IP {
	broadcast := make(net.IP, len(ipNet.IP))
	copy(broadcast, ipNet.IP)

	for i := range broadcast {
		broadcast[i] |= ^ipNet.Mask[i]
	}

	return broadcast
}

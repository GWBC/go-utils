package utils

import "net"

func _incip(ip net.IP) {
	for i := len(ip) - 1; i >= 0; i-- {
		ip[i]++
		if ip[i] > 0 {
			break
		}
	}
}

func GetAllIPV4(n net.IPNet) []net.IP {
	ips := []net.IP{}

	for ip := n.IP.Mask(n.Mask); n.Contains(ip); _incip(ip) {
		ipcopy := make(net.IP, len(ip))
		copy(ipcopy, ip)
		ips = append(ips, ipcopy)
	}

	return ips
}

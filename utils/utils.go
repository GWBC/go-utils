package utils

import (
	"net"
)

func SplitSlice(slice []byte, chunkSize int) [][]byte {
	var chunks [][]byte
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func GetLocalV4Addrs() (map[string]bool, error) {
	ips := map[string]bool{}

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ips, err
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips[ipNet.IP.String()] = true
			}
		}
	}

	return ips, nil
}

func CopyBytes(data []byte) []byte {
	ret := make([]byte, len(data))
	copy(ret, data)

	return ret
}

func CopyIPNet(netAddr net.IPNet) net.IPNet {
	_, n, _ := net.ParseCIDR(netAddr.String())
	return *n
}

func CopyIP(ip net.IP) net.IP {
	copyIP := make([]byte, len(ip))
	copy(copyIP, ip)

	return copyIP
}

func NetaddrToRange(netAddr net.IPNet) (net.IP, net.IP) {
	startIP := netAddr.IP
	endIP := CopyIP(startIP)

	for i := range endIP {
		endIP[i] |= ^netAddr.Mask[i]
	}

	return startIP, endIP
}

func FilterSlice[T any](slice []T, filterFun func(T) bool) []T {
	j := 0
	for _, v := range slice {
		if !filterFun(v) {
			slice[j] = v
			j++
		}
	}

	return slice[:j]
}

package utils

import (
	"net"
	"regexp"
	"strings"
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

func CopyBytes(data []byte) []byte {
	ret := make([]byte, len(data))
	copy(ret, data)

	return ret
}

// 拷贝网络地址
func CopyIPNet(netAddr net.IPNet) net.IPNet {
	_, n, _ := net.ParseCIDR(netAddr.String())
	return *n
}

// 拷贝IP地址
func CopyIP(ip net.IP) net.IP {
	copyIP := make([]byte, len(ip))
	copy(copyIP, ip)

	return copyIP
}

// 获取指定网络的开始和结束IP
func NetaddrToRange(netAddr net.IPNet) (net.IP, net.IP) {
	startIP := netAddr.IP
	endIP := CopyIP(startIP)

	for i := range endIP {
		endIP[i] |= ^netAddr.Mask[i]
	}

	return startIP, endIP
}

// 过滤切片
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

// 获取公网IP
func GetPublicIPv4(addr string, retryCount int) (net.IP, error) {
	var retErr error
	for range retryCount {
		ip, err := _getIP(addr)
		if err != nil {
			retErr = err
			continue
		}

		return ip, err
	}

	return nil, retErr
}

// 获取本机IP
func GetLocalIPv4(cardName string) (map[string]bool, error) {
	cards, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	ips := map[string]bool{}

	for _, card := range cards {
		if len(cardName) != 0 {
			if !strings.EqualFold(cardName, card.Name) {
				continue
			}

			addrs, err := card.Addrs()
			if err != nil {
				return nil, err
			}

			for _, addr := range addrs {
				iaddr := addr.(*net.IPNet)
				v4 := iaddr.IP.To4()
				if v4 == nil {
					continue
				}

				ips[v4.String()] = true
			}

			return ips, nil
		}

		if (card.Flags&net.FlagUp) == 0 || (card.Flags&net.FlagLoopback) == 1 {
			continue
		}

		addrs, err := card.Addrs()
		if err != nil {
			return nil, err
		}

		for _, addr := range addrs {
			iaddr := addr.(*net.IPNet)
			v4 := iaddr.IP.To4()
			if v4 == nil {
				continue
			}

			ips[v4.String()] = true
		}
	}

	return ips, nil
}

// 获取指定网络中的所有IP
func NetToAllIPv4(n net.IPNet) []net.IP {
	ips := []net.IP{}

	for ip := n.IP.Mask(n.Mask); n.Contains(ip); _incip(ip) {
		ipcopy := make(net.IP, len(ip))
		copy(ipcopy, ip)
		ips = append(ips, ipcopy)
	}

	return ips
}

///////////////////////////////////////////////////////

func _incip(ip net.IP) {
	for i := len(ip) - 1; i >= 0; i-- {
		ip[i]++
		if ip[i] > 0 {
			break
		}
	}
}

func _getIP(addr string) ([]byte, error) {
	data, err := Get(addr, nil, nil)
	if err != nil {
		return nil, err
	}

	re, err := regexp.Compile(`\d+\.\d+\.\d+\.\d+`)
	if err != nil {
		return nil, err
	}

	ips := re.FindAllString(string(data), -1)
	for _, ip := range ips {
		netIP := net.ParseIP(ip)
		if netIP != nil {
			return netIP, nil
		}
	}

	return nil, nil
}

package nettun

import (
	"fmt"
	"net"
)

type IPVerson int

const (
	IPV4 IPVerson = iota
	IPV6
	Unknown
)

func ParseIPVersion(data []byte) IPVerson {
	version := data[0] >> 4

	switch version {
	case 4:
		return IPV4
	case 6:
		return IPV6
	}

	return Unknown
}

func ParseIPv4Addr(data []byte) (srcIP, dstIP net.IP, err error) {
	if len(data) < 20 {
		return nil, nil, fmt.Errorf("data len < 20")
	}

	if ParseIPVersion(data) != IPV4 {
		return nil, nil, fmt.Errorf("not ipv4")
	}

	// 提取源IP（12-15字节）和目标IP（16-19字节）
	srcIP = net.IPv4(data[12], data[13], data[14], data[15])
	dstIP = net.IPv4(data[16], data[17], data[18], data[19])

	return srcIP, dstIP, nil
}

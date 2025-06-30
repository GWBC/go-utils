package netset

import (
	"errors"
	"fmt"
	"net"
)

type IPNet struct {
	net.IPNet
}

func (n IPNet) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", n.String())), nil
}

func (n *IPNet) UnmarshalJSON(b []byte) error {
	if len(b) <= 2 {
		return errors.New("json data is error")
	}

	data := b[1 : len(b)-1]
	_, ret, err := net.ParseCIDR(string(data))
	if err != nil {
		return err
	}

	n.IPNet = *ret

	return nil
}

type RouteInfo struct {
	Netaddr    IPNet  //网络地址
	TargetAddr net.IP //目标地址，为nil，则表示链路地址
	Metric     uint32
}

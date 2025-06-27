package netset

import (
	"encoding/json"
	"net"
)

type RouteInfo struct {
	Netaddr    net.IPNet
	TargetAddr net.IP //为nil，则表示链路地址
	Metric     uint32
}

func (r RouteInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"Netaddr":    r.Netaddr.String(),
		"TargetAddr": r.TargetAddr.String(),
		"Metric":     r.Metric,
	})
}

func (ip *RouteInfo) UnmarshalJSON(b []byte) error {
	datas := map[string]any{}
	if err := json.Unmarshal(b, &datas); err != nil {
		return err
	}

	netAddr := datas["Netaddr"]
	_, n, err := net.ParseCIDR(netAddr.(string))
	if err != nil {
		return err
	}

	ip.TargetAddr = net.ParseIP(datas["TargetAddr"].(string))
	ip.Netaddr = *n

	metric := datas["Metric"]
	if metric != nil {
		ip.Metric = uint32(metric.(float64))
	}

	return nil
}

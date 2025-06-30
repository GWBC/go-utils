package netset

import (
	"encoding/json"
	"net"
)

type RouteInfo struct {
	Netaddr    net.IPNet //网络地址
	TargetAddr net.IP    //目标地址，为nil，则表示链路地址
	Metric     uint32
}

func (r RouteInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"Netaddr":    r.Netaddr.String(),
		"TargetAddr": r.TargetAddr.String(),
		"Metric":     r.Metric,
	})
}

func (r *RouteInfo) UnmarshalJSON(b []byte) error {
	datas := map[string]any{}
	if err := json.Unmarshal(b, &datas); err != nil {
		return err
	}

	netAddr := datas["Netaddr"]
	_, n, err := net.ParseCIDR(netAddr.(string))
	if err != nil {
		return err
	}

	r.TargetAddr = net.ParseIP(datas["TargetAddr"].(string))
	r.Netaddr = *n

	metric := datas["Metric"]
	if metric != nil {
		r.Metric = uint32(metric.(float64))
	}

	return nil
}

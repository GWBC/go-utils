//go:build windows

package netset

import (
	"fmt"
	"net"
	"os/exec"
)

const (
	netshCmdTemplateSetAddr  = "netsh interface ip set address name=\"%s\" static %s %s none"
	netshCmdTemplateAddRoute = "netsh interface ipv4 add route %s/%d \"%s\" %s metric=%d store=active"
	netshCmdTemplateDelRoute = "netsh interface ipv4 delete route %s/%d \"%s\" %s store=active"
	netshCmdTemplateSetTMU   = "netsh interface ipv4 set subinterface \"%s\" mtu=%d store=persistent"
)

type Netset struct {
	devName string
}

func (n *Netset) Init(devName string) error {
	n.devName = devName
	return nil
}

func (n *Netset) SetIPAddresses(addr net.IPNet) error {
	mask := net.IP(addr.Mask)
	cmds := []string{}
	cmds = append(cmds, "/C")
	cmds = append(cmds, fmt.Sprintf(netshCmdTemplateSetAddr, n.devName, addr.IP.String(), mask.String()))

	if len(cmds) == 0 {
		return nil
	}

	return runNetsh(cmds)
}

func (n *Netset) AddRoute(route *RouteInfo) error {
	n.DelRoute(route)

	ip := route.Netaddr.IP.String()
	mask, _ := route.Netaddr.Mask.Size()

	cmds := []string{}
	cmds = append(cmds, "/C")
	cmds = append(cmds, fmt.Sprintf(netshCmdTemplateAddRoute,
		ip, mask, n.devName, route.TargetAddr.String(), route.Metric))

	if len(cmds) == 0 {
		return nil
	}

	return runNetsh(cmds)
}

func (n *Netset) DelRoute(route *RouteInfo) error {
	ip := route.Netaddr.IP.String()
	mask, _ := route.Netaddr.Mask.Size()

	cmds := []string{}
	cmds = append(cmds, "/C")
	cmds = append(cmds, fmt.Sprintf(netshCmdTemplateDelRoute,
		ip, mask, n.devName, route.TargetAddr.String()))

	if len(cmds) == 0 {
		return nil
	}

	return runNetsh(cmds)
}

func (n *Netset) SetMTU(mtu int) error {
	cmds := []string{}
	cmds = append(cmds, "/C")
	cmds = append(cmds, fmt.Sprintf(netshCmdTemplateSetTMU, n.devName, mtu))

	if len(cmds) == 0 {
		return nil
	}

	return runNetsh(cmds)
}

func (n *Netset) AddRoutes(routes []RouteInfo) error {
	for _, route := range routes {
		err := n.AddRoute(&route)
		if err != nil {
			return err
		}
	}

	return nil
}

func runNetsh(cmds []string) error {
	cmd := exec.Command("cmd.exe", cmds...)
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

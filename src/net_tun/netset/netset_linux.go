//go:build linux

package netset

import (
	"net"
	"syscall"

	"github.com/vishvananda/netlink"
)

type Netset struct {
	link netlink.Link
}

func (n *Netset) Init(devName string) error {
	link, err := netlink.LinkByName(devName)
	if err != nil {
		return err
	}

	n.link = link
	return netlink.LinkSetUp(n.link)
}

func (n *Netset) SetIPAddresses(addr net.IPNet) error {
	taddr := netlink.Addr{}
	taddr.IPNet = &addr
	taddr.Peer = nil
	taddr.Flags &^= syscall.IFF_POINTOPOINT
	err := netlink.AddrReplace(n.link, &taddr)
	if err != nil {
		return err
	}

	return nil
}

func (n *Netset) AddRoute(route *RouteInfo) error {
	r := &netlink.Route{
		Dst:       &route.Netaddr,
		Priority:  int(route.Metric),
		Scope:     netlink.SCOPE_LINK,
		LinkIndex: n.link.Attrs().Index,
		Protocol:  syscall.RTPROT_STATIC,
		Family:    netlink.FAMILY_V4,
	}

	if route.TargetAddr != nil {
		r.Gw = route.TargetAddr
		r.Scope = netlink.SCOPE_UNIVERSE
	}

	netlink.RouteDel(r)
	return netlink.RouteAdd(r)
}

func (n *Netset) DelRoute(route *RouteInfo) error {
	r := &netlink.Route{
		Dst:       &route.Netaddr,
		Priority:  int(route.Metric),
		Scope:     netlink.SCOPE_LINK,
		LinkIndex: n.link.Attrs().Index,
		Protocol:  syscall.RTPROT_STATIC,
		Family:    netlink.FAMILY_V4,
	}

	if route.TargetAddr != nil {
		r.Gw = route.TargetAddr
		r.Scope = netlink.SCOPE_UNIVERSE
	}

	return netlink.RouteDel(r)
}

func (n *Netset) SetMTU(mtu int) error {
	return netlink.LinkSetMTU(n.link, mtu)
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

package main

import (
	"fmt"
	"net"
	"time"

	"github.com/GWBC/go-utils/utils"
	"github.com/GWBC/go-utils/utils/net_tun/netset"
	//p2plog "github.com/ipfs/go-log/v2"
	// "github.com/libp2p/go-libp2p"
	// "github.com/libp2p/go-libp2p/core/crypto"
	// "github.com/libp2p/go-libp2p/core/host"
	// "github.com/libp2p/go-libp2p/core/network"
	// "github.com/libp2p/go-libp2p/core/peer"
	// "github.com/libp2p/go-libp2p/p2p/protocol/autonatv2"
	// "github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/client"
	// "github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/relay"
)

func NatTest() {
	netaddr := net.IPNet{}
	netaddr.IP = net.ParseIP("10.0.0.23")
	netaddr.Mask = net.IPv4Mask(255, 255, 255, 0)
	netset.StartForward()
	netset.DelNatMasquerade(netaddr)
	netset.AddNatMasquerade(netaddr)

	time.Sleep(1000 * time.Second)
	netset.StopForward()
}

// func newRelay() (host.Host, peer.AddrInfo) {

// 	//可以固化下来
// 	priv, _, err := crypto.GenerateKeyPairWithReader(crypto.RSA, 2048, rand.Reader)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	relaySvr, err := libp2p.New(libp2p.EnableRelay(),
// 		libp2p.Identity(priv),
// 		libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/2222"))
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	fmt.Println("ID", relaySvr.ID(), "Addr", relaySvr.Addrs())

// 	_, err = relay.New(relaySvr, relay.WithInfiniteLimits())
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	return relaySvr, peer.AddrInfo{ID: relaySvr.ID(), Addrs: relaySvr.Addrs()}
// }

// func newHost(relayAddr peer.AddrInfo, port int16) host.Host {
// 	nodeHost, err := libp2p.New(
// 		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port)),
// 		libp2p.EnableRelay(),
// 		libp2p.EnableNATService(),
// 		libp2p.EnableHolePunching())
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	_, err = autonatv2.New(nodeHost)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	err = nodeHost.Connect(context.Background(), relayAddr)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	nodeHost.SetStreamHandler("/test", func(s network.Stream) {
// 		defer s.Close()
// 		_, err := io.Copy(s, s)
// 		if err != nil {
// 			log.Fatal(err.Error())
// 		}
// 	})

// 	return nodeHost
// }

// func TestP2P(){
// 	//p2plog.SetAllLoggers(p2plog.LevelDebug)

// 	r, raddr := newRelay()
// 	nodeA := newHost(raddr, 2230)
// 	nodeB := newHost(raddr, 2231)

// 	p2pAddr, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/p2p/%s/p2p-circuit/p2p/%s", r.ID(), nodeA.ID()))
// 	p2pAddrs := peer.AddrInfo{
// 		ID:    nodeA.ID(),
// 		Addrs: []multiaddr.Multiaddr{p2pAddr},
// 	}

// 	//预约中继服务
// 	_, err := client.Reserve(context.Background(), nodeA, raddr)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	err = nodeB.Connect(context.Background(), peer.AddrInfo{ID: nodeA.ID(), Addrs: nodeA.Addrs()})
// 	if err != nil {
// 		err = nodeB.Connect(context.Background(), p2pAddrs)
// 		if err != nil {
// 			log.Fatal(err.Error())
// 		}
// 	}

// 	s, err := nodeB.NewStream(network.WithAllowLimitedConn(context.Background(), "test"), nodeA.ID(), "/test")
// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	s.Write([]byte("test"))

// 	buf := make([]byte, 4096)

// 	for {
// 		n, err := s.Read(buf)
// 		if err != nil {
// 			return
// 		}

// 		fmt.Println(string(buf[:n]))
// 		time.Sleep(1 * time.Second)
// 		s.Write(buf[:n])

// 		x := nodeB.Network().ConnsToPeer(nodeA.ID())
// 		fmt.Println(x)
// 	}

// 	select {}
// }

func main() {
	//NatTest()

	eth2, _ := utils.GetLocalIPv4("eth2")
	fmt.Println(eth2)
}

package main

import (
	"net"
	"time"

	"github.com/GWBC/go-utils/utils/net_tun/netset"
)

func main() {
	_, netAddr, _ := net.ParseCIDR("10.0.0.0/24")
	netset.StartForward()
	netset.DelNatMasquerade(*netAddr)
	netset.AddNatMasquerade(*netAddr)

	time.Sleep(10 * time.Second)
	netset.StopForward()
}

// package main

// import (
// 	"fmt"
// 	"log"

// 	"github.com/lysShub/divert-go"
// 	"github.com/pkg/errors"
// 	"golang.org/x/sys/windows"
// 	"gvisor.dev/gvisor/pkg/tcpip/header" // go get gvisor.dev/gvisor@go
// )

// var _ = divert.MustLoad(divert.DLL)

// func main() {
// 	d, err := divert.Open("ip.DstAddr>=10.0.0.2 and ip.DstAddr<=10.0.1.255 and !loopback", divert.Network, 0, 0)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var b = make([]byte, 1536)
// 	var addr divert.Address
// 	for {
// 		n, err := d.Recv(b[:cap(b)], &addr)
// 		if err != nil {
// 			if errors.Is(err, windows.ERROR_INSUFFICIENT_BUFFER) {
// 				continue
// 			}
// 			log.Fatal(err)
// 		} else if n == 0 {
// 			continue
// 		}

// 		if !addr.IPv6() {
// 			if n >= header.IPv4MinimumSize+header.TCPMinimumSize {
// 				iphdr := header.IPv4(b[:n])
// 				tcphdr := header.TCP(iphdr[iphdr.HeaderLength():])

// 				fmt.Printf("%s:%d --> %s:%d \n",
// 					iphdr.SourceAddress().String(),
// 					tcphdr.SourcePort(),
// 					iphdr.DestinationAddress().String(),
// 					tcphdr.DestinationPort(),
// 				)
// 				d.Send(b, &addr)
// 			}
// 		} else {
// 			if n >= header.IPv6MinimumSize+header.TCPMinimumSize {
// 				iphdr := header.IPv6(b[:n])
// 				tcphdr := header.TCP(iphdr[header.IPv6MinimumSize:])

// 				fmt.Printf("%s:%d --> %s:%d \n",
// 					iphdr.SourceAddress().String(),
// 					tcphdr.SourcePort(),
// 					iphdr.DestinationAddress().String(),
// 					tcphdr.DestinationPort(),
// 				)
// 			}

// 			d.Send(b, &addr)
// 		}
// 	}
// }

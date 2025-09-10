package main

import (
	"fmt"
	"net"
	"path/filepath"
	"strings"
	"time"

	"github.com/GWBC/go-utils/test"
	"github.com/GWBC/go-utils/utils"
	jsengine "github.com/GWBC/go-utils/utils/js_engine"
	"github.com/GWBC/go-utils/utils/net_tun/netset"
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

func TestJS() {
	fpath := filepath.Join(utils.Pwd(), "..", "test", "test.js")

	js, _ := jsengine.New("js")
	err := js.Require(fpath, "test")
	if err != nil {
		panic(err)
	}

	ret, err := js.RunString("test.Home()")
	if err != nil {
		panic(err)
	}

	fmt.Println(ret.(string))
}

func TestMPD() {
	data := jsengine.BlibiliData2MPD(test.BliData, "/api/videojs/proxy-play?proxy=", jsengine.SelectAudio)
	fmt.Println(data)
}

func TestQbittorrent() {
	api := utils.QbittorrentApi{}
	api.SetHost("192.168.1.10:3457").Login("admin", "")
	data, err := api.GetAllInfo()
	if err == nil {
		fmt.Println(data)
	}

	//api.Add("magnet:?xt=urn:btih:9B20B088B19C98147D8666A545AF401D77250849", "/downloads/xiaoma")

	for k, v := range *data.Torrents {
		if strings.Contains(*v.Name, "小马") {
			api.Stop(k)
			api.Start(k)
			api.Delete(k, true)
		}
	}
}

func main() {
	//NatTest()

	//TestJS()

	//TestMPD()

	TestQbittorrent()
}

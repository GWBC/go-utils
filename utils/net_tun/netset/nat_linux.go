//go:build linux

package netset

import (
	"context"
	"net"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

var natsLock sync.RWMutex
var nats []net.IPNet

func SetWriteVNetFun(writeVnet func([]byte)) {
}

func StartForward() error {
	return os.WriteFile("/proc/sys/net/ipv4/ip_forward", []byte("1"), 0644)
}

func StopForward() {
	for _, nat := range nats {
		DelNatMasquerade(nat)
	}
}

func DelNatMasquerade(netAddr net.IPNet) {
	netAddr.IP = netAddr.IP.Mask(netAddr.Mask)
	output, err := execCmd("nft", "--handle", "list", "chain", "nat", "POSTROUTING")
	if err == nil {
		handles := []string{}
		strNetaddr := netAddr.String()

		outputs := strings.Split(output, "\n")
		for _, line := range outputs {
			if !strings.Contains(line, strNetaddr) {
				continue
			}

			handleIndex := strings.Index(line, "# handle")
			if handleIndex == -1 {
				continue
			}

			handleStr := line[handleIndex+len("# handle"):]
			handleStr = strings.TrimSpace(handleStr)
			handles = append(handles, handleStr)
		}

		for _, handle := range handles {
			execCmd("nft", "delete", "rule", "nat", "POSTROUTING", "handle", handle)
		}
	}

	for {
		_, err := execCmd("iptables", "-t", "nat", "-D", "POSTROUTING",
			"-s", netAddr.String(), "-j", "MASQUERADE")
		if err != nil {
			break
		}
	}

	natsLock.Lock()
	defer natsLock.Unlock()
	for i, nat := range nats {
		if netAddr.String() == nat.String() {
			copy(nats[i:], nats[i+1:])
		}
	}
}

func AddNatMasquerade(netAddr net.IPNet) error {
	netAddr.IP = netAddr.IP.Mask(netAddr.Mask)
	_, err := execCmd("nft", "add", "rule", "ip", "nat", "POSTROUTING",
		"ip", "saddr", netAddr.String(), "counter", "masquerade")

	if err != nil {
		_, err = execCmd("iptables", "-t", "nat", "-A", "POSTROUTING",
			"-s", netAddr.String(), "-j", "MASQUERADE")
		return err
	}

	return nil
}

func execCmd(name string, args ...string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, name, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

//go:build windows

package netset

import (
	"errors"
	"sync"

	"github.com/lysShub/divert-go"
	"golang.org/x/sys/windows"
)

type HookPacketFun = func(handle *divert.Handle, addr *divert.Address, pkg []byte)

type Windivert struct {
	wg     sync.WaitGroup
	handle *divert.Handle
}

func (w *Windivert) Start(filter string, layers divert.Layer, hookFun HookPacketFun) error {
	handle, err := divert.Open(filter, layers, 0, 0)
	if err != nil {
		return err
	}

	w.handle = handle
	w.wg.Add(1)
	go func() {
		defer w.wg.Done()

		var addr divert.Address
		var data = make([]byte, 64*1024)

		for {
			n, err := handle.Recv(data, &addr)
			if err != nil {
				if errors.Is(err, windows.ERROR_INSUFFICIENT_BUFFER) {
					continue
				}

				return
			}

			hookFun(w.handle, &addr, data[:n])
		}
	}()

	return nil
}

func (w *Windivert) Stop() {
	if w.handle != nil {
		w.handle.Close()
	}

	w.wg.Wait()
	w.handle = nil
}

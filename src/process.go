package utils

import (
	"errors"
	"os/exec"
	"strings"

	"github.com/shirou/gopsutil/process"
)

func KillProcess(name string) error {
	processes, _ := process.Processes()
	for _, p := range processes {
		n, _ := p.Name()
		if strings.EqualFold(n, name) ||
			strings.EqualFold(n, name+".exe") {
			err := p.Kill()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func StartProcess(exe string, args ...string) error {
	cmd := exec.Command(exe, args...)
	if cmd == nil {
		return errors.New("start process error")
	}

	err := cmd.Start()
	if err != nil {
		return err
	}

	return nil
}

func ExeIsRun(name string) bool {
	processes, _ := process.Processes()
	for _, p := range processes {
		n, _ := p.Name()
		if strings.EqualFold(n, name) ||
			strings.EqualFold(n, name+".exe") {
			return true
		}
	}

	return false
}

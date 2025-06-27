package ulog

import (
	"fmt"
)

type Console struct {
}

func (c *Console) Write(data []byte) (int, error) {
	fmt.Print(string(data))
	return len(data), nil
}

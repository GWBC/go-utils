package ulog

import (
	"github.com/sirupsen/logrus"
)

func NewHook(fun WriteFun) *Hook {
	h := &Hook{
		WriteFun: fun,
	}

	return h
}

//////////////////////////////////////////////////////////

type WriteFun func(entry *logrus.Entry) error

type Hook struct {
	WriteFun WriteFun
}

func (d *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (d *Hook) Fire(entry *logrus.Entry) error {
	return d.WriteFun(entry)
}

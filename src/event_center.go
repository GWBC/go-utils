package utils

import (
	"sync"
)

type HandlerInterface interface {
	Proc(arg any)
}

type EventCenter struct {
	lock   sync.Mutex
	events map[any][]HandlerInterface
}

func (e *EventCenter) Init() {
	e.events = make(map[any][]HandlerInterface)
}

func (e *EventCenter) UnInit() {

}

func (e *EventCenter) On(eName any, handler HandlerInterface) {
	e.lock.Lock()
	defer e.lock.Unlock()

	q := e.events[eName]
	if q == nil {
		q = []HandlerInterface{handler}
		e.events[eName] = q
	} else {
		e.events[eName] = append(q, handler)
	}
}

func (e *EventCenter) Off(eName any, handler HandlerInterface) {
	e.lock.Lock()
	defer e.lock.Unlock()
	delete(e.events, eName)
}

func (e *EventCenter) Emit(eName any, arg any) {
	e.lock.Lock()
	defer e.lock.Unlock()

	hs := e.events[eName]
	if hs == nil {
		return
	}

	go func() {
		for _, v := range hs {
			v.Proc(arg)
		}
	}()
}

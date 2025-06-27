package utils

import "sync"

type Single[T any] struct {
	once sync.Once
	obj  *T
}

func (s *Single[T]) Instance(newFun ...func() *T) *T {
	s.once.Do(func() {
		if len(newFun) == 0 {
			s.obj = new(T)
		} else {
			s.obj = newFun[0]()
		}
	})

	return s.obj
}

package utils

import "sync"

type Coroutine func(args ...any)

func GoRun(wg *sync.WaitGroup, fun Coroutine, args ...any) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		fun(args)
	}()
}

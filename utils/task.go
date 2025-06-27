package utils

import (
	"context"
	"errors"
	"sync"
	"time"
)

type TaskResult[T any] func(result *T)

type TaskInterface[T any] interface {
	Proc() *T
}

type TaskQueue[T any] struct {
	maxIdleTime time.Duration
	cmdLock     sync.Mutex
	cmds        map[string]*SafeQueue
	ctx         context.Context
	cancel      context.CancelFunc
}

func (t *TaskQueue[T]) Init(maxIdleTime time.Duration) {
	t.cmds = make(map[string]*SafeQueue)
	t.ctx, t.cancel = context.WithCancel(context.Background())

	if maxIdleTime == 0 {
		t.maxIdleTime = time.Hour
	} else {
		t.maxIdleTime = maxIdleTime
	}
}

func (t *TaskQueue[T]) UnInit() {
	if t.ctx == nil {
		return
	}

	select {
	case <-t.ctx.Done():
		return
	default:
		t.cancel()
		for _, v := range t.cmds {
			v.UnInit()
		}
	}
}

func (t *TaskQueue[T]) PushTask(queueID string, cmd TaskInterface[T], resultFun TaskResult[T]) {
	t.getQueue(queueID).Push(cmd, func(result any) {
		resultFun(result.(*T))
	})
}

func (t *TaskQueue[T]) PushTaskAndWaitResult(queueID string, cmd TaskInterface[T]) *T {
	return t.getQueue(queueID).PushAndWaitResult(cmd).(*T)
}

func (t *TaskQueue[T]) getQueue(id string) *SafeQueue {
	t.cmdLock.Lock()
	defer t.cmdLock.Unlock()
	queue, ok := t.cmds[id]
	if ok {
		return queue
	}

	queue = &SafeQueue{}
	queue.Init(t.ctx, t.maxIdleTime)
	queue.Run(func(data any) any {
		cmd, ok := data.(TaskInterface[T])
		if ok {
			return cmd.Proc()
		}

		return errors.New("data is not CmdInterface")
	}, func() bool {
		t.cmdLock.Lock()
		defer t.cmdLock.Unlock()
		delete(t.cmds, id)
		return false
	})

	return queue
}

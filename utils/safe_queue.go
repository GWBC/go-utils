package utils

import (
	"container/list"
	"context"
	"sync"
	"time"
)

type DataProcFun func(data any) any
type DataTimeoutFun func() bool
type ProcResultFun func(result any)

type DataInfo struct {
	data   any
	result any
}

type SafeQueue struct {
	lock      sync.Mutex
	queue     *list.List
	writeChan chan bool
	timer     *time.Timer
	timeout   time.Duration

	wg     sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
}

func (s *SafeQueue) Init(ctx context.Context, timeout time.Duration) {
	s.queue = list.New()
	s.writeChan = make(chan bool, 20)
	s.ctx, s.cancel = context.WithCancel(ctx)

	s.timeout = timeout

	if timeout > 0 {
		s.timer = time.NewTimer(s.timeout)
	} else {
		s.timer = time.NewTimer(time.Hour)
		s.timer.Stop()
	}
}

func (s *SafeQueue) UnInit() {
	select {
	case <-s.ctx.Done():
		return
	default:
		s.cancel()
		s.wg.Wait()
	}
}

func (s *SafeQueue) Push(data any, resultFun ProcResultFun) {
	tdata := &DataInfo{
		data:   data,
		result: resultFun,
	}

	s.lock.Lock()
	defer s.lock.Unlock()
	s.queue.PushBack(tdata)
	s.writeChan <- true
}

func (s *SafeQueue) PushAndWaitResult(data any) any {
	tdata := &DataInfo{
		data:   data,
		result: make(chan any),
	}

	s.lock.Lock()
	s.queue.PushBack(tdata)
	s.writeChan <- true
	s.lock.Unlock()

	return <-tdata.result.(chan any)
}

func (s *SafeQueue) Pop() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.queue.Remove(s.queue.Back())
}

func (s *SafeQueue) Run(procFun DataProcFun, timeoutFun DataTimeoutFun) {
	GoRun(&s.wg, func(args ...any) {
	task_exit:
		for {
			select {
			case <-s.ctx.Done():
				break task_exit
			case <-s.writeChan:
				data := s.getData()
				if data != nil {
					if retChan, ok := data.result.(chan any); ok {
						ret := procFun(data.data)
						retChan <- ret
					} else if retFun, ok := data.result.(ProcResultFun); ok {
						ret := procFun(data.data)
						retFun(ret)
					} else {
						panic("safe queue struct error")
					}
				}
			case <-s.timer.C:
				if timeoutFun != nil {
					if !timeoutFun() {
						break task_exit
					}
				}
			}

			s.timer.Reset(s.timeout)
		}

		//释放定时器
		if !s.timer.Stop() {
			<-s.timer.C
		}
	})
}

func (s *SafeQueue) getData() *DataInfo {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.queue.Len() == 0 {
		return nil
	}

	data := s.queue.Back()
	s.queue.Remove(data)
	return data.Value.(*DataInfo)
}

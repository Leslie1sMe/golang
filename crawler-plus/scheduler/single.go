package scheduler

import "go_code/crawler-plus/engine"

type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

func (s SimpleScheduler) ConfigWorkerChan(ch chan engine.Request) {
	s.WorkerChan = ch
}

func (s SimpleScheduler) Submit(request engine.Request) {
	go func() { s.WorkerChan <- request }()
}

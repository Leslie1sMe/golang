package scheduler

import "go_code/crawler-plus/engine"

//QueuedScheduler
type QueuedScheduler struct {
	RequestChannel chan engine.Request
	WorkerChannel  chan chan engine.Request
}

func (q *QueuedScheduler) ConfigWorkerChan(request chan engine.Request) {
	panic("implement me")
}

//Submit
func (q *QueuedScheduler) Submit(request engine.Request) {
	q.RequestChannel <- request
}

//WorkChanFree
func (q *QueuedScheduler) WorkChanFree(worker chan engine.Request) {
	q.WorkerChannel <- worker
}

//Begin
func (q *QueuedScheduler) Begin() {
	q.RequestChannel = make(chan engine.Request)
	q.WorkerChannel = make(chan chan engine.Request)
	go func() {
		var requestQueue []engine.Request
		var workerQueue []chan engine.Request
		for {
			var freeRequest engine.Request
			var freeWorker chan engine.Request
			if len(workerQueue) > 0 && len(requestQueue) > 0 {
				freeWorker = workerQueue[0]
				freeRequest = requestQueue[0]
			}
			select {
			case request := <-q.RequestChannel:
				requestQueue = append(requestQueue, request)

			case worker := <-q.WorkerChannel:
				workerQueue = append(workerQueue, worker)
			case freeWorker <- freeRequest:
				workerQueue = workerQueue[1:]
				requestQueue = requestQueue[1:]
			}
		}
	}()
}

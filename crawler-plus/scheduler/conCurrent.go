package scheduler

import (
	"go_code/crawler-plus/engine"
)

//ConScheduler
type ConScheduler struct {
	WorkerChannel chan engine.Request
}

//Submit 把接收到的请求都传送到channel里去
func (c *ConScheduler) Submit(request engine.Request) {
	go func() { c.WorkerChannel <- request }()
}

//ConfigWorkerChan
func (c *ConScheduler) ConfigWorkerChan(ch chan engine.Request) {
	c.WorkerChannel = ch
}

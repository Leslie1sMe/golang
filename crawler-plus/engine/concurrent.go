package engine

type Scheduler interface {
	Submit(Request)
	ConfigWorkerChan(chan Request)
}

//ConcurrentEngine
type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

//Run
func (c *ConcurrentEngine) Run(seed ...Request) {
	var in = make(chan Request)
	var out = make(chan ParseResult)
	c.Scheduler.ConfigWorkerChan(in)
	for i := 0; i < c.WorkerCount; i++ {
		CreateWorkers(in, out)
	}
	for _, request := range seed {
		c.Scheduler.Submit(request)
	}
	for {
		results := <-out
		for _, request := range results.Requests {
			c.Scheduler.Submit(request)
		}

	}
}

//CreateWorkers
func CreateWorkers(in chan Request, out chan ParseResult) {
	go func() {
		for {
			requests := <-in
			result, err := Worker(requests)
			if err != nil {
				continue
			}
			out <- result
		}

	}()
}

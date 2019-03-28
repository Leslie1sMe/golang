package engine

//ConcurrentEngine
type ConcurrentEngine struct {
	Scheduler   Scheduler
	Fetcher     Fetcher
	WorkerCount int
}

//Run
func (c *ConcurrentEngine) Run(seed ...Request) {
	var out = make(chan ParseResult)
	c.Scheduler.Begin()
	for i := 0; i < c.WorkerCount; i++ {
		CreateWorkers(out, c.Scheduler, c.Fetcher)
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
func CreateWorkers(out chan ParseResult, s Scheduler, f Fetcher) {
	in := make(chan Request)
	go func() {
		for {
			s.WorkChanFree(in)
			requests := <-in
			result, err := f.Work(requests)
			if err != nil {
				continue
			}
			out <- result
		}

	}()
}

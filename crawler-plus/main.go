package main

import (
	"go_code/crawler-plus/engine"
	"go_code/crawler-plus/parsers"
	"go_code/crawler-plus/scheduler"
)

func main() {
	var requests = engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parsers.GetCitiesList,
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.ConScheduler{},
		WorkerCount: 100,
	}
	e.Run(requests)
}

package main

import (
	"go_code/crawler-plus/config"
	"go_code/crawler-plus/engine"
	"go_code/crawler-plus/fetcher"
	"go_code/crawler-plus/parsers"
	"go_code/crawler-plus/scheduler"
	"go_code/crawler-plus/writer"
)

func main() {
	var requests = engine.Request{
		Url:        config.Fetch_url,
		ParserFunc: parsers.GetCitiesList,
	}

	dataChan := writer.MongoDbWriter{}.Write("lzl", "profile") //?待优化
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		Fetcher:   &fetcher.ForgedFetcher{}, //两种爬取方式，支持简单和并发爬取
		Writer: engine.WriteWorker{
			Payload: dataChan,               //将要存取的数据放入channel,解耦分离给存储模块存储
			Storage: writer.MongoDbWriter{}, //三种写入方式支持mongodb，redis，elastic Api
		},
		WorkerCount: 100,
	}
	e.Run(requests)
}

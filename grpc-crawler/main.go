package main

import (
	"go_code/crawler-plus/config"
	"go_code/crawler-plus/engine"
	"go_code/crawler-plus/fetcher"
	"go_code/crawler-plus/parsers"
	"go_code/crawler-plus/scheduler"
	"grpc-crawler/client"
)

func main() {
	var requests = engine.Request{
		Url:        config.Fetch_url,
		ParserFunc: parsers.GetCitiesList,
	}
	dataChan := client.GrpcService{}.GrpcWrite(":9002")
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		Fetcher:   &fetcher.ForgedFetcher{}, //两种爬取方式，支持简单和并发爬取
		Writer: engine.WriteWorker{
			Payload: dataChan,             //将要存取的数据放入channel,解耦分离给存储模块存储
			GSaver:  client.GrpcService{}, //实现RPC模式,将数据发送给远端服务器存储
		},
		WorkerCount: 100,
	}
	e.Run(requests)
}

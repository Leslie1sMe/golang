package main

import (
	"github.com/monnand/goredis"
	"github.com/olivere/elastic"
	"go_code/crawler-plus/engine"
	"go_code/crawler-plus/fetcher"
	"go_code/crawler-plus/parsers"
	"go_code/crawler-plus/scheduler"
	"go_code/crawler-plus/writer"
	"log"
)

func main() {
	var requests = engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parsers.GetCitiesList,
	}
	elkClient, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		log.Fatal(err)
	}
	redisClient := &goredis.Client{Addr: ":6379", Db: 1}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		Fetcher:     &fetcher.ForgedFetcher{Client: redisClient},
		Writer:      &writer.ElkWriter{Client: elkClient},
		WorkerCount: 100,
	}
	e.Run(requests)
}

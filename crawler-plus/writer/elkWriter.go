package writer

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"log"
)

type ElkWriter struct {
}

func (e ElkWriter) Write(args ...string) chan interface{} {
	dataChan := make(chan interface{})

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			item := <-dataChan
			resp, err := client.Index().Index(args[0]).Type(args[1]).BodyJson(item).Do(context.Background())
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(resp)
		}
	}()
	return dataChan
}

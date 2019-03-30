package main

import (
	"github.com/olivere/elastic"
	"go_code/crawler-rpc/elkWriteService"
	"go_code/crawler-rpc/rpc-drive/server"
	"log"
)

type ItemSaverService struct {
	Index  string
	Type   string
	Client *elastic.Client
}

func main() {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		log.Fatal(err)
	}
	server.RpcServer(":9002", elkWriteRpc.ElkWriteService{
		Client: client,
		Index:  "za",
		Type:   "profile",
	})
}

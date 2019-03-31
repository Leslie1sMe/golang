package main

import (
	"fmt"
	"github.com/olivere/elastic"
	"go_code/crawler-rpc/rpc-drive/server"
	"go_code/crawler-rpc/writeService"
	"log"
)

func main() {
	go ServerRpc(":9002")
}

func ServerRpc(host string) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		log.Fatal(err)
	}
	err = rpc_drive.RpcServer(host, writeServiceRpc.WriteServiceRpc{
		Client: client,
		Index:  "za",
		Type:   "profile",
	})
	fmt.Println(err)
}

package main

import (
	"fmt"
	"github.com/olivere/elastic"
	"go_code/crawler-rpc/rpc-drive"
	"go_code/crawler-rpc/writeService/client"
	"log"
)

func main() {
	ServerRpc(":9002")
}

//ServerRpc
func ServerRpc(host string) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		log.Fatal(err)
	}
	err = rpc_drive.RpcServer(host, &writeServiceRpc.WriteServiceRpc{
		Client: client,
		Index:  "za",
		Type:   "profile",
	})
	fmt.Println(err)
}

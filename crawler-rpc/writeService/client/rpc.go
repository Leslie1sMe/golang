package rpc_client

import (
	"fmt"
	"go_code/crawler-rpc/rpc-drive/client"
)

func Write(host string, item interface{}) {
	client, err := rpc_drive.NewRpcClient(host)
	if err != nil {
		fmt.Println(err)
	}
	res := new(string)
	err = client.Call("WriteService.RpcWriter", item, &res)
	if err != nil {
		fmt.Println(err)
	}
}

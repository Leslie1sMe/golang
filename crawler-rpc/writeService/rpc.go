package writeService

import (
	"fmt"
	"go_code/crawler-rpc/rpc-drive"
)

func Write(host string) chan interface{} {
	dataChan := make(chan interface{})
	var reply = ""
	client, err := rpc_drive.NewRpcClient(host)
	if err != nil {
		fmt.Println(err)
	}
	go func() {
		for {
			item := <-dataChan
			err = client.Call("WriteServiceRpc.Write", item, &reply)
			if err != nil {
				fmt.Println(err)
			}
		}
	}()
	return dataChan
}

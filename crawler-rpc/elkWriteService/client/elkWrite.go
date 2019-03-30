package client

import (
	"fmt"
	"go_code/crawler-plus/model"
	"go_code/crawler-rpc/rpc-drive/client"
)

func ElkWriter(items model.Profile, result *string) error {
	client, err := client.NewRpcClient(":9002")
	if err != nil {
		return err
	}
	err = client.Call("ElkWriteService.Write", items, &result)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

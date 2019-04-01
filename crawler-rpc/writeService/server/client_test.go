package main

import (
	"fmt"
	"go_code/crawler-rpc/rpc-drive"
	"go_code/crawler/model"
	"testing"
)

func TestServerRpc(t *testing.T) {
	go ServerRpc(":9002")

	client, err := rpc_drive.NewRpcClient(":9002")
	if err != nil {
		t.Error(err)
	}
	result := ""
	item := model.Profile{
		Name:              "Les",
		Province:          "Beijing",
		Age:               "27",
		MarriedOrNot:      "yes",
		Height:            "181",
		Income:            "30k",
		DegreeOfEducation: "Master",
	}
	err = client.Call("WriteService.RpcWriter", item, &result)
	if err != nil {
		fmt.Println(err)
	}
}

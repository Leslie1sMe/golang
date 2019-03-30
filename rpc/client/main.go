package main

import (
	"fmt"
	"go_code/rpc"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	Conn, err := net.Dial("tcp", ":9001")
	if err != nil {
		fmt.Println(err)
	}
	client := jsonrpc.NewClient(Conn)
	var result float64
	err = client.Call("CrawlerRpcService.Add", crawler_rpc.Args{A: 10, B: 20}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

package main

import (
	"fmt"
	"go_code/rpc"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	err := rpc.Register(crawler_rpc.CrawlerRpcService{})
	if err != nil {
		fmt.Println(err)
	}
	listener, err := net.Listen("tcp", ":9001")
	if err != nil {
		fmt.Println(err)
	}
	for {
		Conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("等待连接...")
		go jsonrpc.ServeConn(Conn)

	}
}

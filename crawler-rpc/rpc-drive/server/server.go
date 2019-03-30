package server

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func RpcServer(host string, service interface{}) error {
	rpc.Register(service)
	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		jsonrpc.ServeConn(conn)
	}
}

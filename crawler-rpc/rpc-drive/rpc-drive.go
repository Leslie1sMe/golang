package rpc_drive

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//RpcServer
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

//NewRpcClient
func NewRpcClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(conn), nil
}

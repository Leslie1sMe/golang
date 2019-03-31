package rpc_drive

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func NewRpcClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(conn), nil
}

package grpc_support

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

//GrpcServer
func GrpcServer(host string) (listener net.Listener, s *grpc.Server) {
	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("listen on %s\n", host)
	s = grpc.NewServer()
	return listener, s
}

//NewClient
func NewClient(host string) (conn *grpc.ClientConn, error error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		error = err
		return
	}
	return conn, nil
}

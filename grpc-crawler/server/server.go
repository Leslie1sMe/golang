package main

import (
	"fmt"
	"go_code/grpc-crawler/etcd"
	"go_code/grpc-crawler/grpc-support"
	pb "go_code/grpc-crawler/proto"
	"golang.org/x/net/context"
)

type WriteService struct {
}

func (w *WriteService) GrpcWrite(ctx context.Context, req *pb.ClientRequest) (resp *pb.ClientResponse, err error) {
	profile := req.GetItem()
	if profile != nil {
		fmt.Println(profile)
	}
	resp = &pb.ClientResponse{
		Reply: "ok",
	}
	return resp, nil
}

func main() {

	port := "http://127.0.0.1:2379"
	etcd.RegisterEtcdService([]string{port})
	l, s := grpc_support.GrpcServer("127.0.0.1:2379")
	var u = &WriteService{}
	pb.RegisterWriteServiceServer(s, u)
	s.Serve(l)

}

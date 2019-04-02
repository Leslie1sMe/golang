package main

import (
	"fmt"
	"golang.org/x/net/context"
	"grpc-crawler/grpc-support"
	pb "grpc-crawler/proto"
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
	port := ":9002"
	l, s := grpc_support.GrpcServer(port)
	var u = &WriteService{}
	pb.RegisterWriteServiceServer(s, u)
	s.Serve(l)

}

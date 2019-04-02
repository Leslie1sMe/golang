package client

import (
	"fmt"
	"golang.org/x/net/context"
	"grpc-crawler/grpc-support"
	pb "grpc-crawler/proto"
)

type GrpcService struct {
}

func (g GrpcService) GrpcWrite(host string) chan pb.Profile {
	dataChan := make(chan pb.Profile)
	conn, err := grpc_support.NewClient(host)
	if err != nil {
		fmt.Println(err)
	}
	client := pb.NewWriteServiceClient(conn)
	req := new(pb.ClientRequest)
	req.Index = "za"
	req.ElkType = "profile"
	go func() {
		for {
			item := <-dataChan
			fmt.Println(item)
			fmt.Printf("item的类型是%T\n", item)
			req.Item = &item
			req.ElkType = "profile"
			req.Index = "za"
			resp, err := client.GrpcWrite(context.Background(), req)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(resp)

		}
	}()
	return dataChan
}

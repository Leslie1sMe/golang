package client

import (
	"fmt"
	"go_code/grpc-crawler/etcd"
	"go_code/grpc-crawler/grpc-support"
	pb "go_code/grpc-crawler/proto"
	"golang.org/x/net/context"
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

func (g GrpcService) EtcdWrite(name string, host string, endpoint []string) chan pb.Profile {
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
			etcd.RegisterClientService(name, "host", endpoint)
		}

	}()
	return dataChan
}

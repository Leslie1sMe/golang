package etcd

import (
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/client"
	"golang.org/x/net/context"
	"log"
	"os"
	"time"
)

type Client interface {
	HeartBeat()
}

type ClientService struct {
	KeyApi client.KeysAPI
	Name   string
	Ip     string
}

type ClientInfo struct {
	Name string
	Ip   string
	Pid  int
}

func RegisterClientService(name, ip string, endpoints []string) {
	cfg := client.Config{
		Endpoints:               endpoints,
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal("Error: cannot connec to etcd:", err)
	}
	s := &ClientService{
		Name:   name,
		KeyApi: client.NewKeysAPI(c),
		Ip:     ip,
	}
	go s.HeartBeat()
}

func (c *ClientService) HeartBeat() {
	api := c.KeyApi
	for {
		clientInfo := &ClientInfo{
			Name: c.Name,
			Ip:   c.Ip,
			Pid:  os.Getpid(),
		}
		key := "workers/" + c.Name
		v, err := json.Marshal(clientInfo)
		if err != nil {
			log.Print("json marshal failed:%s", err)
			break
		}
		resp, err := api.Set(context.Background(), key, string(v), &client.SetOptions{
			TTL: 10 * time.Second,
		})
		if err != nil {
			log.Print("api configure failed:%s", err)
			break
		} else {
			fmt.Println(resp.Node.Value)
		}
		time.Sleep(5 * time.Second)

	}
}

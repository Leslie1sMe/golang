package main

import (
	"fmt"
	"github.com/monnand/goredis"
	"log"
	"time"
)

type Message struct {
}

func main() {
	//redis基本操作
	var client goredis.Client
	//设置基本redis信息
	client.Addr = "127.0.0.1:6379"
	client.Db = 1
	//string操作
	client.Set("test", []byte("hello"))
	res, _ := client.Get("string")
	fmt.Println(string(res))
	client.Del("string")
	//list操作
	list := []string{"a", "b", "c", "d"}
	for _, v := range list {
		client.Rpush("list", []byte(v))
	}
	listRes, error := client.Lrange("list", 0, 3)
	if error != nil {
		log.Fatal(error)
	}
	for i, v := range listRes {
		fmt.Println(i, ":", string(v))
	}
	client.Del("list")
	//Publish/Subscribe操作
	sub := make(chan string, 1)
	sub <- "foo"
	messages := make(chan goredis.Message, 0)
	go client.Subscribe(sub, nil, nil, nil, messages)

	time.Sleep(10 * 1000 * 1000)
	client.Publish("foo", []byte("bar"))

	msg := <-messages
	println("received from:", msg.Channel, " message:", string(msg.Message))

	close(sub)
	close(messages)

}

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"ume/common"
)

//server端分发信息
var conOnline = make(map[string]net.Conn)

//信息channel
var MessageQueue = make(chan string, 1000)
var quitChan = make(chan bool)

func Listen(conn net.Conn) {
	mes := make([]byte, 1024)
	for {
		num, error := conn.Read(mes)
		if num != 0 {
			message := string(mes[0:num])
			MessageQueue <- message
		}
		if error != nil {
			if error == io.EOF {
				continue
			} else {
				panic(error)
			}
		}

	}
}

func Consumer() {
	for {
		select {
		case message := <-MessageQueue:
			ProcessMessage(message)
		case <-quitChan:
			break
		}
	}
}

func ProcessMessage(message string) {
	content := strings.Split(message, "#")

	if len(content) > 1 {
		address := content[0]
		messageSent := content[1]
		addr := strings.Trim(address, " ")
		if conn, ok := conOnline[addr]; ok {
			_, error := conn.Write([]byte(messageSent))
			if error != nil {
				log.Fatal(error)
			}
		}
		fmt.Println(message)
	}

}

func main() {
	//初始化一个监听listener
	listenSocket, error := net.Listen("tcp", "127.0.0.1:7788")
	common.CheckError(error)
	defer listenSocket.Close()
	fmt.Println("聊天机器人开始工作")
	//select模式for循环拉取chan里的消息
	go Consumer()
	for {
		conn, error := listenSocket.Accept()
		if error != nil {
			panic(error)
		}
		conOnline[conn.RemoteAddr().String()] = conn
		for ip, _ := range conOnline {
			fmt.Println(ip)
		}
		go Listen(conn)
	}

}

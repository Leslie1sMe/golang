package main

import (
	"fmt"
	"io"
	"net"
	"ume/common"
)

var MessageQueue = make(chan string)
var quietChan = make(chan bool)

func Listen(conn net.Conn) {
	mes := make([]byte, 1024)
	for {
		num, error := conn.Read(mes)
		if num != 0 {
			message := string(mes[0:num])
			fmt.Println(message)
			//	MessageQueue <- message
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
			fmt.Println(message)
		}
	}
}

func ProcessMessage(message string) {
	fmt.Println(message)
}

func main() {
	listenSocket, error := net.Listen("tcp", "127.0.0.1:7788")
	common.CheckError(error)
	defer listenSocket.Close()
	fmt.Println("聊天机器人开始工作")
	//go Consumer()
	for {
		conn, error := listenSocket.Accept()
		if error != nil {
			panic(error)
		}
		go Listen(conn)
	}

}

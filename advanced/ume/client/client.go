package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
	"ume/common"
)

func SendMessage(d net.Conn) {
	var input string
	for {
		Reader := bufio.NewReader(os.Stdin)
		line, _, _ := Reader.ReadLine()
		input = string(line)
		if strings.ToUpper(input) == "EXIT" {
			d.Close()
			break
		}
		fmt.Println(time.Now().Format("2006-01-02 15:04:05") + "发送成功")

		_, err := d.Write([]byte(input))
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	common.LogError()
	d, err := net.Dial("tcp", "127.0.0.1:7788")
	if err != nil {
		common.CheckError(err)
	}
	fmt.Printf("您(%s)已经连接到聊天服务器%s\n", d.LocalAddr().String(), d.RemoteAddr().String())
	defer d.Close()
	go SendMessage(d)
	buf := make([]byte, 1024)
	for {
		num, error := d.Read(buf)
		common.CheckError(error)
		fmt.Printf("收到消息:\n%v", string(buf[0:num]))
	}

}

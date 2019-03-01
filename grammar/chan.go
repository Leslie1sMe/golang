package main

import (
	"fmt"
	"time"
)

func test_channel(ch chan int) {
	ch <- 1
	fmt.Scanln("this is goroutine1")
	time.Sleep(time.Second)
}

func main() {
	ch := make(chan int, 1)
	go test_channel(ch)
	fmt.Println("this is goroutine2")
	time.Sleep(2 * time.Second)

}

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter   int64
	WaitGroup sync.WaitGroup
)

func main() {

	WaitGroup.Add(2)
	go test(1)
	go test(2)
	WaitGroup.Wait()
	fmt.Println("Final Counter", counter)
}

func test(a int) {
	defer WaitGroup.Done()
	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}

}

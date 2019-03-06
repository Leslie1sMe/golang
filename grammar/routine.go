package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func main() {
	//for i := 0; i < 1000; i++ {
	//	go func(i int) {
	//		fmt.Println("goroutine", i)
	//		//runtime.Gosched()
	//	}(i)
	//
	//}
	//
	//time.Sleep(time.Minute)
	//WaitGroup.Add(2)
	//go test(1)
	//go test(2)
	//WaitGroup.Wait()
	//fmt.Println("Final Counter", counter)
	//使用调度器

	go func() {
		for x := 1; x < 100; x++ {
			if x%10 == 0 {
				runtime.Gosched() //使用Gosched()主动让出执行权
			}
			fmt.Println("routine1:" + strconv.Itoa(x))
		}
	}()

	go func() {
		for x := 100; x < 200; x++ {
			fmt.Println("routine2:" + strconv.Itoa(x))
		}
	}()

	time.Sleep(2 * time.Second)

}

//func test(a int) {
//defer WaitGroup.Done()
//for count := 0; count < 2; count++ {
//	atomic.AddInt64(&counter, 1)
//	runtime.Gosched()
//}

//}

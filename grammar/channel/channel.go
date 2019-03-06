package main

var ch chan int

func chanDemo() {
	//d := make(chan int)
	//c := make(chan int, 10)
	//go func() {
	//	for {
	//		s := <-d
	//		fmt.Println(s)
	//	}
	//}()
	//go func() {
	//	for {
	//		n := <-d
	//		fmt.Println(n)
	//	}
	//}()
	//d <- 3
	//d <- 4
	//d <- 6
	//time.Sleep(time.Millisecond)
	//ch <- 1
	//ch <- 1
	//fmt.Println("这是goroutine1")

}

func main() {
	//ch = make(chan int)
	//go func(ch chan int) {
	//	ch <- 1
	//}(ch)
	//time.Sleep(2 * time.Second)
	////select语句，轮询查找符合的case
	//select {
	//case ch <- 1:
	//	fmt.Println("正在写哦")
	//case <-ch:
	//	fmt.Println("正在读哦")
	//default:
	//	fmt.Println("noting to do")
	//}

	//select超时应用

	//ch = make(chan int)
	//timeout := make(chan bool, 1)
	//go func() {
	//	ch <- 1
	//	time.Sleep(time.Second)
	//	timeout <- true
	//}()
	//
	//select {
	//case <-ch:
	//	fmt.Println("读取ch中...")
	//case <-timeout:
	//	fmt.Println("超时啦")
	//
	//}

	//ch = make(chan int)
	//go func() {
	//	for x := 1; x < 100; x++ {
	//		if x == 10 {
	//			<-ch
	//			//runtime.Gosched() //使用Gosched()主动让出执行权
	//		}
	//		fmt.Println("routine1:" + strconv.Itoa(x))
	//	}
	//}()
	//
	//go func() {
	//	for x := 100; x < 200; x++ {
	//		fmt.Println("routine2:" + strconv.Itoa(x))
	//	}
	//
	//	ch <- 1
	//}()
	//
	//time.Sleep(2 * time.Second)

}

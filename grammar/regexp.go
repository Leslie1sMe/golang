package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
)

func main() {
	Work()
}

func Work() {
	ch := make(chan int)
	for i := 0; i <= 0; i++ {
		go Spy(i, ch)
	}
	for i := 0; i <= 0; i++ {
		fmt.Printf("第%d页爬取完成\n", <-ch)
	}

}

func Spy(i int, ch chan int) {
	url := "http://tieba.baidu.com/f?kw=张国荣&ie=utf-8&pn=" + strconv.Itoa(50*i)
	res, err := Get(url, i)
	if err != nil {
		return
	}
	reg := regexp.MustCompile(` (?s:(.*?))`)
	fmt.Println(res)
	result := reg.FindAllStringSubmatch(res, -1)
	for _, v := range result {
		fmt.Println(v[1])
	}
	ch <- i
}

func Get(url string, i int) (result string, err error) {
	resp, error := http.Get(url)
	if error != nil {
		err = error
		return
	}
	defer resp.Body.Close()
	buf := make([]byte, 4096)
	for {
		n, error := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println(strconv.Itoa(i) + "读取完成")
			break
		}
		if error != nil && error != io.EOF {
			err = error
			return
		}
		result += string(buf[:n])

	}
	return
}

package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Info struct {
	Title       string
	Description string
}

var ch chan int

func main() {
	var page = 1
	fmt.Println("正在爬取中...")
	BeginToWork(page, "科比")
}

func BeginToWork(n int, name string) {
	for i := 0; i <= n; i++ {
		url := "http://tieba.baidu.com/f?kw=" + name + "&ie=utf-8&pn=" + strconv.Itoa(50*i)
		html, error := SpyWork(url)
		if error != nil {
			fmt.Println(error)
		}
		HandleImages(html)
	}

}

func HandleHtml(html string) {
	a := make([]string, 10)
	b := make([]string, 10)
	var c []Info
	regTitle := regexp.MustCompile(`class="j_th_tit ">(?s:(.*?))</a>`)
	titleSlice := regTitle.FindAllStringSubmatch(html, -1)
	regDescription := regexp.MustCompile(`<div class="threadlist_abs threadlist_abs_onlyline ">(?s:(.*?))</div>`)
	descSlice := regDescription.FindAllStringSubmatch(html, -1)
	for _, v := range titleSlice {
		a = append(a, v[1])
	}
	for _, v := range descSlice {
		b = append(b, v[1])
	}
	for i := 0; i <= len(titleSlice); i++ {
		c = append(c, Info{Title: a[i], Description: b[i]})
	}
	fmt.Println(c)

}

func HandleImages(html string) {
	regImg := regexp.MustCompile(`data-original="(?s:(.*?))"`)
	imgSlice := regImg.FindAllStringSubmatch(html, -1)
	for _, v := range imgSlice {
		str, err := SpyWork(v[1])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(v[1] + "爬取成功\n")
		path := "/Users/leslie/go/src/go_code/images/" + fmt.Sprintf("%10v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000)) + ".jpg"
		file, _ := os.Create(path)
		file.Write([]byte(str))
		defer file.Close()
	}
}

func SpyWork(url string) (html string, error error) {
	resp, err := http.Get(url)
	if err != nil {
		error = err
		return
	}
	defer resp.Body.Close()
	var res = make([]byte, 4096)
	for {
		n, err := resp.Body.Read(res)
		if n == 0 {
			fmt.Println("Successfully downloaded")
			fmt.Println(ch)
			break
		}
		if err != nil && err != io.EOF {
			error = err
			return
		}
		html += string(res[:n])
	}
	return
}

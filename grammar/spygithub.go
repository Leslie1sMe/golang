package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
	"regexp"
)

type Github struct {
	Name        string
	Description string
}

func main() {
	fmt.Println("请输入要爬取的github用户名")
	var name string
	fmt.Scanln(&name)
	handle(name)
}

func handle(name string) {
	if result, error := work(name); error != nil {
		fmt.Println(error)
	} else {
		var a, b []string
		var c []Github
		reg1 := regexp.MustCompile(` <span class="repo js-pinnable-item" title="(?s:(.*?))">(?s:(.*?))</span>`)
		NameSlice1 := reg1.FindAllStringSubmatch(result, -1)
		for _, v := range NameSlice1 {
			a = append(a, v[1])
		}
		reg2 := regexp.MustCompile(` <p class="pinned-item-desc text-gray text-small d-block mt-2 mb-3">(?s:(.*?))</p>`)
		NameSlice2 := reg2.FindAllStringSubmatch(result, -1)
		for _, v := range NameSlice2 {
			b = append(b, v[1])
		}
		if len(a) == len(b) {
			for i := 0; i < len(a); i++ {
				d := Github{Name: a[i], Description: b[i]}
				c = append(c, d)
			}
			fmt.Println(c)
			//然后就可以存mysql了

		}
	}
}
func work(name string) (result string, error error) {
	url := "https://github.com/" + name
	if resp, err := http.Get(url); err != nil {
		error = err
		return
	} else {
		defer resp.Body.Close()
		buf := make([]byte, 4096)
		for {
			n, err := resp.Body.Read(buf)
			if n == 0 {
				fmt.Println("读取完成")
				break
			}
			if err != nil && err != io.EOF {
				error = err
				return
			}
			result += string(buf[:n])
		}
	}
	return

}

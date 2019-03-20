package spy

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

//开始工作
func GetToWork(url string, compileString string) {
	var ch = make(chan int)
	html, err := Spy(url)
	if err != nil {
		fmt.Println(err)
	}
	imgList := GetImagesList(html, compileString)
	for index, v := range imgList {
		go ImgHandler(index, ch, v)
	}
	for i := 1; i <= len(imgList); i++ {
		fmt.Printf("第%d张爬取完成\n", <-ch)
	}
}

//获取目标地址的html代码
func Spy(url string) (html string, err error) {
	resp, getError := http.Get(url)
	if getError != nil {
		err = getError
		return
	}
	defer resp.Body.Close()
	res := make([]byte, 1024)
	for {
		num, getError := resp.Body.Read(res)
		if getError != nil && getError != io.EOF {
			err = getError
			return
		}
		if num == 0 {
			break
		}
		html += string(res[:num])
	}
	return
}

//根据传入compileString,解析出目标图片集url
func GetImagesList(html string, compileString string) []string {
	reg := regexp.MustCompile(compileString)
	var s []string
	objSlice := reg.FindAllStringSubmatch(html, -1)
	for _, v := range objSlice {
		s = append(s, v[1])
	}
	fmt.Println("数量:", len(objSlice))
	return s
}

//处理图片地址并保存在本地
func ImgHandler(num int, ch chan int, url string) (error error) {
	res, err := Spy(url)
	if err != nil {
		error = err
		return
	}
	filePath := "/Users/leslie/go/src/go_code/images/spy" + strconv.Itoa(num) + ".jpg"
	file, err := os.Create(filePath)
	defer file.Close()
	if err != nil {
		error = err
		return
	}
	file.Write([]byte(res))
	ch <- (num + 1)
	return
}

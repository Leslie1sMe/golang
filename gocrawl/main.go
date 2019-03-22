package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

type Profile struct {
	Name   string
	Age    string
	Weight string
	Height string
	Image  string
}

func main() {
	var ch = make(chan int)
	for i := 1; i <= 133; i++ {
		go HandlePage(i, ch)
	}

	for i := 1; i <= 1330; i++ {
		fmt.Printf("第%d条插入elk完成\n", <-ch)
	}
}

//解析每页数据
func HandlePage(i int, ch chan int) {
	url := "https://vip.jiayuan.com/cjjllist.php?&p=" + strconv.Itoa(i)
	html := ParseHtml(url)
	profile := GetToWork(html)
	var num = 0
	for _, v := range profile {
		num++
		PutIntoElk(ch, num, v)
	}
}

//解析url
func ParseHtml(url string) (html string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
	}
	if html, err := ioutil.ReadAll(resp.Body); err == nil {
		return string(html)
	} else {
		panic(err)
	}
	return
}

//开始工作
func GetToWork(html string) (profile []Profile) {
	nameReg := regexp.MustCompile(`class="memberName">(?s:(.*?))</a></p>`)
	nameSlice := nameReg.FindAllStringSubmatch(html, -1)
	linkReg := regexp.MustCompile(`<p><a target="_blank" href="(?s:(.*?))" class="memberName">`)
	linkSlice := linkReg.FindAllStringSubmatch(html, -1)
	var profiles []Profile
	for index, v := range nameSlice {
		link := "https://vip.jiayuan.com" + linkSlice[index][1]
		profileHtml := ParseHtml(link)
		ageReg := regexp.MustCompile(`<span class="item">年龄：(?s:(.*?))</span>`)
		heightReg := regexp.MustCompile(`<span class="item">身高：(?s:(.*?))</span>`)
		weightReg := regexp.MustCompile(`<span class="item">体重：(?s:(.*?))</span>`)
		imgReg := regexp.MustCompile(`<img src="(?s:(.*?))" alt="" />`)
		ageSlice := ageReg.FindAllStringSubmatch(profileHtml, -1)
		heightSlice := heightReg.FindAllStringSubmatch(profileHtml, -1)
		weightSlice := weightReg.FindAllStringSubmatch(profileHtml, -1)
		imgSlice := imgReg.FindAllStringSubmatch(profileHtml, -1)
		profiles = append(profiles, Profile{Name: v[1], Age: ageSlice[0][1], Weight: weightSlice[0][1], Height: heightSlice[0][1], Image: imgSlice[0][1]})
	}
	return profiles
}

//存入ELK中
func PutIntoElk(ch chan int, num int, profile Profile) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		fmt.Println(err)
	}
	_, err = client.Index().Index("shijijiayuan").Type("profile").BodyJson(profile).Do(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	ch <- num
}

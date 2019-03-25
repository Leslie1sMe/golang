package main

import (
	"fmt"
	"go_code/crawler/model"
	"go_code/ip_pools"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func main() {

	webUrl := "http://www.zhenai.com/zhenghun"
	proxy, client := ip_pools.SetProxy()
	html, error := getHtml(proxy, client, webUrl)
	if error != nil {
		fmt.Println(error)
	} else {
		var num int
		ch := make(chan int)
		go func() {
			//获取城市列表
			cityList := CityListParser(html)
			for _, v := range cityList {
				html, error = getHtml(proxy, client, v)
				if error != nil {
					fmt.Println(error)
					fmt.Printf("切换ip\n", proxy)
					proxy, client = ip_pools.SetProxy()
					html, error = getHtml(proxy, client, v)
					continue
				}
				//获取用户列表
				userList := UserListParser(html)
				for _, v := range userList {
					num++
					html, error = getHtml(proxy, client, v)
					if error != nil {
						fmt.Println(error)
						fmt.Printf("切换ip\n", proxy)
						proxy, client = ip_pools.SetProxy()
						html, error = getHtml(proxy, client, v)
						continue
					}
					//获取用户信息列表
					userProfileList := UserProfileParser(num, ch, html)
					for _, v := range userProfileList {
						fmt.Println(v)
					}

				}
			}
		}()
		for i := 0; i < 10000; i++ {
			fmt.Println(<-ch)
		}
	}
}

func getHtml(proxy *url.URL, client *http.Client, webUrl string) (html []byte, err error) {

	request, err := http.NewRequest("GET", webUrl, nil)
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set("User-Agent", ip_pools.GetAgent())
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	request.Header.Set("Connection", "keep-alive")
	response, err := client.Do(request)
	if err != nil || response.StatusCode != 200 {
		fmt.Printf("遇到了错误,并切换ip %s\n", err)
		proxy, _ := ip_pools.SetProxy()
		fmt.Println("切换ip成功\n", proxy)
	}
	html, error := ioutil.ReadAll(response.Body)
	if error != nil {
		return nil, error
	}
	return html, nil
}

//城市列表解析器
func CityListParser(content []byte) []string {
	var cityList []string
	reg := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[1-9a-z]+)" data-v-5e16505f>([^<]+)</a>`)
	s := reg.FindAllStringSubmatch(string(content), -1)
	for _, v := range s {
		cityList = append(cityList, v[1])
	}
	return cityList
}

//用户列表解析器
func UserListParser(content []byte) []string {
	var userList []string
	reg := regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]+)</a>`)
	s := reg.FindAllStringSubmatch(string(content), -1)
	for _, v := range s {
		userList = append(userList, v[1])
	}
	return userList
}

//用户资料解析器
func UserProfileParser(num int, ch chan int, content []byte) []model.Profile {
	var userProfile []model.Profile
	regNickName := regexp.MustCompile(`<h1 class="nickName" data-v-5b109fc3>([^<]+)</h1>`)
	nameSlice := regNickName.FindAllStringSubmatch(string(content), 1)
	reg := regexp.MustCompile(`<div class="des f-cl" data-v-3c42fade>(?s:(.*?))</div>`)
	s := reg.FindAllStringSubmatch(string(content), 1)
	for _, v := range s {
		trimString := strings.TrimSpace(v[1])
		stringSlice := strings.Split(trimString, "|")
		userProfile = append(userProfile, model.Profile{Name: nameSlice[0][1], Province: stringSlice[0], Age: stringSlice[1], MarriedOrNot: stringSlice[3], Height: stringSlice[4], Income: stringSlice[5], DegreeOfEducation: stringSlice[2]})
	}
	ch <- num
	return userProfile
}

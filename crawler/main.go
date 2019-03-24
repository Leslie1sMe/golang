package main

import (
	"fmt"
	"go_code/crawler/ip_pools"
	"go_code/crawler/model"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {

	res, err := ip_pools.GetPosition()
	if err != nil {
		fmt.Println(err)
	}
	ip_pools.SaveIpToRedis(res)
	ip_pools.GetIp()
	s := ip_pools.GetAgent()
	//url := "http://www.zhenai.com/zhenghun"
	//html, error := getHtml(url)
	//if error != nil {
	//	fmt.Println(error)
	//} else {
	//	cityList := printCityList(html)
	//	for _, v := range cityList {
	//		html, error = getHtml(v)
	//		if error != nil {
	//			fmt.Println(error)
	//			continue
	//		}
	//		userList := printUserList(html)
	//		fmt.Println(userList)
	//		for _, v := range userList {
	//			html, error = getHtml(v)
	//			if error != nil {
	//				fmt.Println(error)
	//				continue
	//			}
	//			//res := printUserProfile(html)
	//			//fmt.Println(string(html))
	//
	//		}
	//	}
	//}
}

func getHtml(url string) (html []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.StatusCode)
	}
	html, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}
	return html, nil
}

func printCityList(content []byte) []string {
	var cityList []string
	reg := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[1-9a-z]+)" data-v-5e16505f>([^<]+)</a>`)
	s := reg.FindAllStringSubmatch(string(content), -1)
	for _, v := range s {
		cityList = append(cityList, v[1])
	}
	return cityList
}

func printUserList(content []byte) []string {
	var userList []string
	reg := regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]+)</a>`)
	s := reg.FindAllStringSubmatch(string(content), -1)
	for _, v := range s {
		userList = append(userList, v[1])
	}
	return userList
}

func printUserProfile(content []byte) []model.Profile {
	var profileList []model.Profile
	reg := regexp.MustCompile(`div class="m-btn purple" data-v-bff6f798>([^<]+)</div>`)
	s := reg.FindAllStringSubmatch(string(content), -1)
	for _, v := range s {
		fmt.Println(v)
		//profileList = append(profileList, v[1])
	}
	return profileList
}

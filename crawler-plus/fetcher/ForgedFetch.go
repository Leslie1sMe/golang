package fetcher

import (
	"fmt"
	"github.com/monnand/goredis"
	"go_code/crawler-plus/engine"
	"io/ioutil"
	"ip_pools/agents"
	"net/http"
	"net/url"
	"time"
)

var client = http.Client{}

type ForgedFetcher struct {
	Client goredis.Client
}

func (f *ForgedFetcher) GetProxy() (string, error) {
	if res, err := f.Client.Spop("ip_pools"); err != nil {
		return "", err
	} else {
		return string(res), nil
	}
}

func (f *ForgedFetcher) Fetch(weburl string) (html []byte, error error) {
	request, err := http.NewRequest("GET", weburl, nil)
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set("User-Agent", agents.GetAgent())
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	request.Header.Set("Connection", "keep-alive")
	response, err := client.Do(request)
	if err != nil || response.StatusCode != 200 {
		//fmt.Println("切换ip\n")
		//proxy, err := f.GetProxy()
		//if err != nil {
		//	error = err
		//	return
		//}
		//res, err := url.Parse(proxy)
		//if err != nil {
		//	error = err
		//	return
		//}
		//client = NewClient(res)
		//response, err := client.Do(request)
		//defer response.Body.Close()
		//html, err = ioutil.ReadAll(response.Body)
		if err != nil {
			error = err
			return
		}
	}
	defer response.Body.Close()
	html, err = ioutil.ReadAll(response.Body)
	if err != nil {
		error = err
		return
	}
	return html, nil
}

func (f *ForgedFetcher) Work(request engine.Request) (engine.ParseResult, error) {
	content, err := f.Fetch(request.Url)
	if err != nil {
		return engine.ParseResult{}, err
	}
	return request.ParserFunc(content), nil
}

func NewClient(proxy *url.URL) (client http.Client) {
	client = http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
		Timeout: time.Duration(10 * time.Second),
	}
	return
}

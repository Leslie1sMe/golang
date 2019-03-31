package fetcher

import (
	"github.com/monnand/goredis"
	"go_code/crawler-plus/engine"
	"io/ioutil"
	"net/http"
)

type SimpleFetch struct {
	client goredis.Client
}

func (s SimpleFetch) GetProxy() (string, error) {
	if res, err := s.client.Spop("ip_pools"); err != nil {
		return "", err
	} else {
		return string(res), nil
	}
}

//Work
func (s SimpleFetch) Work(request engine.Request) (engine.ParseResult, error) {
	content, err := s.Fetch(request.Url)
	if err != nil {
		return engine.ParseResult{}, err
	}
	return request.ParserFunc(content), nil
}

//Fetcher
func (s SimpleFetch) Fetch(url string) (html []byte, error error) {
	resp, err := http.Get(url)
	if err != nil {
		error = err
		return
	}
	html, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		error = err
		return
	}
	return html, err
}

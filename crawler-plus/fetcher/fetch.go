package fetcher

import (
	"fmt"
	"io"
	"io/ioutil"
	"ip_pools/agents"
	"net/http"
)

var client http.Client

func Fetch(url string) ([]byte, error) {
	var reader io.Reader
	request, err := http.NewRequest("GET", url, reader)
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set("User-Agent", agents.GetAgent())
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	request.Header.Set("Connection", "keep-alive")
	response, err := client.Do(request)
	if err != nil || response.StatusCode != 200 {
		return nil, err
	}
	defer response.Body.Close()
	html, error := ioutil.ReadAll(response.Body)
	if error != nil {
		return nil, error
	}
	return html, nil
}

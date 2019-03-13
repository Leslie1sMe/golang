package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://github.com/Leslie1sMe"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	sHtml, _ := ioutil.ReadAll(resp.Body)
}

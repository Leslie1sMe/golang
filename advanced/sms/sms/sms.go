package sms

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type SendParam struct {
	ApiKey string
	UseSsl bool
}

var Data map[string]string

//发送短信
func (s *SendParam) Send(phone string, message string) (*http.Response, error) {
	values := url.Values{}
	values.Add("mobile", phone)
	values.Add("message", message)
	query := values.Encode()
	request, err := http.NewRequest("POST", "http://sms-api.luosimao.com/v1/send.json", strings.NewReader(query))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Cache-Control", "no-cache")
	request.SetBasicAuth("api", s.ApiKey)
	fmt.Println(request)
	client := http.DefaultClient
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fmt.Println(resp)
	return resp, err
}

//群发短信
func (s *SendParam) SendBatch(mobileList []string, message string) (*http.Response, error) {
	stringRes := strings.Join(mobileList, ",")
	fmt.Println(stringRes)
	values := url.Values{}
	values.Add("mobile_list", stringRes)
	values.Add("message", message)
	query := values.Encode()
	request, err := http.NewRequest("POST", "https://sms-api.luosimao.com/v1/send_batch.json", strings.NewReader(query))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Cache-Control", "no-cache")
	request.SetBasicAuth("api", s.ApiKey)
	fmt.Println(request)
	client := http.DefaultClient
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fmt.Println(resp)
	return resp, err
}

//查询余额
func (s *SendParam) CheckDeposit() (string, error) {
	request, err := http.NewRequest("GET", "http://sms-api.luosimao.com/v1/status.json", nil)
	if err != nil {
		return "", err
	}
	request.SetBasicAuth("api", s.ApiKey)
	fmt.Println(request)
	client := http.DefaultClient
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), nil
}

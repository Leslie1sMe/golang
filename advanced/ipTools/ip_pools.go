package ip_pools

//实现IP代理池,伪造User-Agent,并返回客户端实例,用于爬虫抓包和科学上网
import (
	"fmt"
	"github.com/monnand/goredis"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

//伪造并返回agent
func GetAgent() string {
	var Agent = [...]string{
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:50.0) Gecko/20100101 Firefox/50.0",
		"Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; en) Presto/2.8.131 Version/11.11",
		"Opera/9.80 (Windows NT 6.1; U; en) Presto/2.8.131 Version/11.11",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; 360SE)",
		"Mozilla/5.0 (Windows NT 6.1; rv:2.0.1) Gecko/20100101 Firefox/4.0.1",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; The World)",
		"User-Agent,Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_8; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		"User-Agent, Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1; Maxthon 2.0)",
		"User-Agent,Mozilla/5.0 (Windows; U; Windows NT 6.1; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	len := len(Agent)
	return Agent[r.Intn(len)]
}

//获取ip源地址IP列表
func GetPosition() (html []byte, err error) {
	url := "https://cn-proxy.com"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	html, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return html, nil
}

//将获取到的ip存入本地redis
func SaveIpToRedis(content []byte) {
	reg := regexp.MustCompile(`<td>([0-9]+\.[0-9]+\.[0-9]+\.[0-9]+)</td>
<td>([0-9]+)</td>`)
	ipSlice := reg.FindAllStringSubmatch(string(content), -1)
	var client goredis.Client
	client.Addr = "127.0.0.1:6379"
	client.Db = 1
	for _, v := range ipSlice {
		s := "http://" + v[1] + ":" + v[2]
		client.Sadd("ip_pools", []byte(s))
	}
}

//随机取redis中的ip进行下一步代理
func GetIp() string {
	var client goredis.Client
	client.Addr = "127.0.0.1:6379"
	client.Db = 1
	res, err := client.Spop("ip_pools")
	if err != nil {
		fmt.Println(err)
	}
	return string(res)
}

//设置proxy并返回client实例
func SetProxy() (proxy *url.URL, client *http.Client) {
	proxy, err := url.Parse(GetIp())
	if err != nil {
		fmt.Println(err)
	}
	client = &http.Client{}
	client = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
		Timeout: time.Duration(20 * time.Second),
	}
	return proxy, client
}

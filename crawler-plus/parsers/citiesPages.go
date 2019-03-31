package parsers

import (
	"fmt"
	"go_code/crawler-plus/engine"
	"regexp"
)

const citiesPage = `(<li class="paging-item"><a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)/([0-9]+))">([^<]+)</a> <!----></li>`

func GetCitiesPages(content []byte) engine.ParseResult {
	re := regexp.MustCompile(citiesPage)
	cities := re.FindAllStringSubmatch(string(content), -1)
	var res = engine.ParseResult{}
	for _, v := range cities {
		fmt.Printf("正在爬取%s\n", v[2])
		res.Requests = append(res.Requests, engine.Request{
			Url:        v[2],
			ParserFunc: GetUserLists,
		})
	}
	return res
}

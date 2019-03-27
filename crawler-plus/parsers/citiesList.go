package parsers

import (
	"go_code/crawler-plus/engine"
	"regexp"
)

const citiesList = `<a href="(http://www.zhenai.com/zhenghun/[1-9a-z]+)" data-v-5e16505f>([^<]+)</a>`

func GetCitiesList(content []byte) engine.ParseResult {
	re := regexp.MustCompile(citiesList)
	cities := re.FindAllStringSubmatch(string(content), -1)
	var res = engine.ParseResult{}
	for _, v := range cities {
		res.Items = append(res.Items, v[2])
		res.Requests = append(res.Requests, engine.Request{
			Url:        v[1],
			ParserFunc: GetUserLists,
		})
	}
	return res
}

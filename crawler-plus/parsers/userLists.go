package parsers

import (
	"go_code/crawler-plus/engine"
	"regexp"
)

const userList = `<a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]+)</a>`

func GetUserLists(content []byte) engine.ParseResult {
	var res = engine.ParseResult{}
	reg := regexp.MustCompile(userList)
	users := reg.FindAllStringSubmatch(string(content), -1)
	for _, v := range users {
		name := v[2]
		res.Items = append(res.Items, name)
		res.Requests = append(res.Requests, engine.Request{
			Url: v[1],
			ParserFunc: func(content []byte) engine.ParseResult {
				return GetUserProfile(content, name)
			},
		})
	}
	return res

}

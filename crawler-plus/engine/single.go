package engine

import (
	"go_code/crawler-plus/fetcher"
)

type SingleEngine struct {
	Scheduler Scheduler
}

//Run
func (s SingleEngine) Run(seed ...Request) {
	var requests []Request
	for _, r := range seed {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parserResult, err := Worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parserResult.Requests...)
	}
}

//Worker
func Worker(request Request) (ParseResult, error) {
	content, err := fetcher.Fetch(request.Url)
	if err != nil {
		return ParseResult{}, err
	}
	return request.ParserFunc(content), nil
}

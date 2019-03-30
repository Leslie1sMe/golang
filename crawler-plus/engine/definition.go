package engine

//Request
type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

//ParseResult
type ParseResult struct {
	Requests []Request
	Items    []interface{}
}

//Scheduler
type Scheduler interface {
	Submit(Request)
	ConfigWorkerChan(chan Request)
	WorkChanFree(chan Request)
	Begin()
}

//Fetcher
type Fetcher interface {
	Fetch(string) ([]byte, error)
	Work(Request) (ParseResult, error)
	GetProxy() (string, error)
}

//Writer
type Writer interface {
	Write(interface{}, string, string) error
}

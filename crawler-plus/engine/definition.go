package engine

import pb "grpc-crawler/proto"

//Request
type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

//ParseResult
type ParseResult struct {
	Requests []Request
	Items    []pb.Profile
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
type Saver interface {
	Write(...string) chan interface{}
}

//WriteWorker
type WriteWorker struct {
	Payload  chan pb.Profile
	Storage  Saver
	RpcSaver RpcSaver
	GSaver   GSaver
}

//RpcSaver
type RpcSaver interface {
	RpcWrite(interface{}, *string) error
}

//GrpcSaver
type GSaver interface {
	GrpcWrite(host string) chan pb.Profile
}

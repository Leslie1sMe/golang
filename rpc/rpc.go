package crawler_rpc

import "github.com/pkg/errors"

type CrawlerRpcService struct {
}

type ElkWriteService struct {
}
type Args struct {
	A, B int
}

func (c CrawlerRpcService) Add(args Args, result *float64) error {
	if args.A == 0 {
		return errors.New("出错啦")
	}
	*result = float64(args.B) + float64(args.A)
	return nil
}

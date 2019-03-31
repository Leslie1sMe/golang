package writer

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
)

type ElkWriter struct {
	Client *elastic.Client
}

//Write
func (e *ElkWriter) Write(item interface{}, args ...string) error {
	resp, err := e.Client.Index().Index(args[0]).Type(args[1]).BodyJson(item).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(resp)
	return nil
}

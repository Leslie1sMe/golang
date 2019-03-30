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
func (e *ElkWriter) Write(item interface{}, index string, elasticType string) error {
	_, err := e.Client.Index().Index(index).Type(elasticType).BodyJson(item).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

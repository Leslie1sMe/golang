package elkWriteRpc

import (
	"github.com/olivere/elastic"
	"go_code/crawler-plus/model"
	"go_code/crawler-plus/writer"
)

type ElkWriteService struct {
	Index  string
	Type   string
	Client *elastic.Client
}

func (i ElkWriteService) Write(items interface{}, result *string) error {
	err := writer.Write(items, i.Index, i.Type)
	if err != nil {
		return err
	}
	*result = "ok"
	return nil
}

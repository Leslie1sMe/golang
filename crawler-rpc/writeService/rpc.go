package writeServiceRpc

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
)

type WriteServiceRpc struct {
	Index  string
	Type   string
	Client *elastic.Client
}

func (w *WriteServiceRpc) RpcWriter(item interface{}, result *string) error {
	indexResponse, err := w.Client.Index().Index(w.Index).Type(w.Type).BodyJson(item).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(indexResponse)
	*result = "ok"
	return nil
}

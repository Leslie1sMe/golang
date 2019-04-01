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

func (w *WriteServiceRpc) RpcWrite(item interface{}, reply *string) error {
	resp, err := w.Client.Index().Index(w.Index).Type(w.Type).BodyJson(item).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(resp)
	*reply = "ok"
	return nil
}

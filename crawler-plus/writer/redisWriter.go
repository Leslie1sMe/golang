package writer

import (
	"encoding/json"
	"fmt"
	"github.com/monnand/goredis"
)

type RedisWriter struct {
	Client *goredis.Client
	Args   string
}

func (r RedisWriter) Write(items interface{}, args ...string) error {
	data, err := json.Marshal(items)
	if err != nil {
		return err
	}
	_, err = r.Client.Hset(args[0], args[1], data)
	if err != nil {
		return err
	} else {
		fmt.Println("插入成功!")
	}
	return nil
}

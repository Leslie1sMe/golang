package writer

import (
	"encoding/json"
	"fmt"
	"github.com/monnand/goredis"
)

type RedisWriter struct {
}

var redis = goredis.Client{Addr: ":6379", Db: 0}

func (r RedisWriter) Write(args ...string) chan interface{} {
	dataChan := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			itemCount++
			items := <-dataChan
			data, err := json.Marshal(items)
			if err != nil {
				fmt.Println(err)
				continue
			}
			_, err = redis.Hset(args[0], string(itemCount), data)
			if err != nil {
				fmt.Println(err)
				continue
			} else {
				fmt.Println("插入成功!")
			}
		}
	}()
	return dataChan

}

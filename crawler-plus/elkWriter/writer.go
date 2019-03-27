package elkWriter

import (
	"context"
	"github.com/olivere/elastic"
	"go_code/crawler-plus/model"
)

func Write(profile model.Profile) (*elastic.IndexResponse, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	resp, err := client.Index().Index("za").Type("profile").BodyJson(profile).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return resp, nil
}

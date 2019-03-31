package writer

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoDbWriter struct {
}

func (m MongoDbWriter) Write(item interface{}, args ...string) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}
	collection := client.Database(args[0]).Collection(args[1])
	res, err := collection.InsertOne(ctx, item)
	if err != nil {
		return err
	}
	id := res.InsertedID
	fmt.Println(id)
	return nil
}

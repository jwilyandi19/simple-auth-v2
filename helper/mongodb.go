package helper

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func OpenMongoDB(ctx context.Context, conf Config) (*mongo.Database, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI(conf.DBAddress)
	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println(conf)
	db := client.Database(conf.DBName)

	return db, nil
}

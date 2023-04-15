package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateMongoInstance(URI string) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func OpenMongoCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = (*mongo.Collection)(client.Database("cluster0").Collection(collectionName))

	return collection
}

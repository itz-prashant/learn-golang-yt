package db

import (
	"context"
	"fmt"
	"time"

	"github.com/itz-prashant/notes-api/internal/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Connect(cfg config.Config) (*mongo.Client, *mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(cfg.MongoUri)
	client, err := mongo.Connect(clientOpts)

	if err != nil {
		return nil, nil, fmt.Errorf("Mongo connection fail")
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, nil, fmt.Errorf("Mongo Ping fail")
	}

	database := client.Database(cfg.MongoDb)

	return client, database, nil
}

func DisConnect(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	return  client.Disconnect(ctx)
}

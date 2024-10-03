package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewMongoClient creates a new MongoDB client
func NewMongoClient(uri string) (*mongo.Client, error) {

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Optionally, you can ping the database to verify the connection
	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, err
	}

	return client, nil
}

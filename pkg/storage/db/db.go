package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client - Database Object
type Client struct {
	*mongo.Client
}

// SetupDB - Connect to Database and return connection
func SetupDB(config string) (*Client, error) {
	clientOptions := options.Client().ApplyURI(config)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	// Check the connection - Reduces Client resiliance...
	// err = client.Ping(context.TODO(), nil)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	return &Client{client}, nil
}

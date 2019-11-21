package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	ticker "github.com/undeadops/giles"
)

func (client *Client) getCollection() *mongo.Collection {
	// Hardsetting DB and collection here.. rework later
	return client.Database("giles").Collection("tickers")
}

// GetAllTickers - Get All Stock Ticker Symbols
func (client *Client) GetAllTickers() ([]*ticker.Ticker, error) {
	// Here's an array in which you can store the decoded documents
	var results []*ticker.Ticker

	// Pass these options to the Find method
	findOptions := options.Find()
	findOptions.SetLimit(20)

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := client.getCollection().Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var tic ticker.Ticker
		err := cur.Decode(&tic)
		if err != nil {
			return nil, err
		}

		results = append(results, &tic)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)
	return results, nil
}

// SaveTicker - Save Stock Ticker Symbol
func (client *Client) SaveTicker(tic *ticker.Ticker) error {
	insertResult, err := client.getCollection().InsertOne(context.TODO(), tic)
	if err != nil {
		return err
	}
	// Future - Raise this to be logged with the event request
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return nil
}

// GetTicker - Get Stock Ticker Symbol
func (client *Client) GetTicker(tic string) (*ticker.Ticker, error) {
	filter := bson.D{{"sysmbol", tic}}

	// create a value into which the result can be decoded
	var result *ticker.Ticker

	err := client.getCollection().FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Found a single document: %+v\n", result)
	return result, nil
}

// DeleteTicker - Delete Stock Ticker Symbol
func (client *Client) DeleteTicker(tic string) error {
	filter := bson.D{{"sysmbol", tic}}

	_, err := client.getCollection().DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	fmt.Println("Connected to MongoDB!")

	// Insert simple data
	collection := client.Database("simpledb").Collection("items")
	item := map[string]interface{}{"name": "Sample Item", "done": false}
	result, err := collection.InsertOne(context.TODO(), item)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted item with ID:", result.InsertedID)
}

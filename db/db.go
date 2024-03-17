package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func init() {
	mongoUrl := ""
	clientOptions := options.Client().ApplyURI(mongoUrl)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}
	if err := client.Ping(context.Background(), nil); err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")
	mongoClient = client
}

func Connect() *mongo.Client {
	return mongoClient
}

func GetCollection(collectionName string) *mongo.Collection {
	if mongoClient == nil {
		panic("MongoDB client is nil. Connection might not have been established.")
	}

	err := mongoClient.Ping(context.Background(), nil)
	if err != nil {
		panic("Failed to ping MongoDB server. Connection might have been lost.")
	}

	database := mongoClient.Database("basic-crud")
	if database == nil {
		panic("Failed to get database 'basic-crud'.")
	}

	collection := database.Collection(collectionName)
	if collection == nil {
		panic(fmt.Sprintf("Failed to get collection '%s'.", collectionName))
	}

	return collection
}

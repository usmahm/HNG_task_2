package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {
	uri := os.Getenv("MONGODB_URI")
	fmt.Println("uri", uri)
	if uri == "" {
		log.Fatal("MONGODB_URI not found in the environment, please add.")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	return client
}

var MongoClient *mongo.Client // To be assigned in main.go

func OpenCollection(collection_name string) *mongo.Collection {
	var collection *mongo.Collection = MongoClient.Database("hnqTask2").Collection(collection_name)

	return collection
}

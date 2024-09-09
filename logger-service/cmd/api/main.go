package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	webPort  = "80"
	rpcPort  = "5001"
	mongoURL  = "mongodb://mongo:27017"
	gRpcPort = "50051"
)

var client *mongo.Client

type Config struct {
}

func main() {
	// connect to mongo
	mongoCLient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}

	client = mongoCLient
}

func connectToMongo() (*mongo.Client, error) {
	// create connection options
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	// connect to mongo
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error connecting to mongodb:", err)
		return nil, err
	}

	return client, nil
}

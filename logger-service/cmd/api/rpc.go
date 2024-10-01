package main

import (
	"context"
	"log"
	"time"

	"github.com/filipe1309/ud-go-microservices/logger-service/data"
)

type RPCServer struct {}

type RPCPayload struct {
	Name string
	Data string
}

func (r *RPCServer) LogInfo(payload RPCPayload, reply *string) error {
	collection := client.Database("logs").Collection("logs")
	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name: payload.Name,
		Data: payload.Data,
		CreatedAt: time.Now(),
	})
	if err != nil {
		log.Println("error writing to mongo", err)
		return err
	}

	*reply = "Processed payload via RPC: " + payload.Name

	return nil
}

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"

	"github.com/filipe1309/ud-go-microservices/logger-service/data"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	webPort  = "80"
	rpcPort  = "5001"
	mongoURL = "mongodb://mongo:27017"
	gRpcPort = "50051"
)

var client *mongo.Client

type Config struct {
	Models data.Models
}

func main() {
	// connect to mongo
	mongoCLient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}

	client = mongoCLient

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// close connection
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Panic("Error disconnecting from mongodb:", err)
		}
	}()

	app := Config{
		Models: data.New(client),
	}

	// start rpc server
	// Register RPC server
	rpcServer := new(RPCServer)
	err = rpc.Register(rpcServer)
	if err != nil {
		log.Panic("Error registering rpc server:", err)
	}
	go app.rpcListen()

	// start web server
	// go app.serve()

	log.Println("Starting logger service on port", webPort)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic("Error on server ListenAndServe:", err)
	}
}

func (app *Config) serve() {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic("Error on server ListenAndServe:", err)
	}
}

func (app *Config) rpcListen() error {
	log.Println("Starting rpc service on port", rpcPort)
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", rpcPort))
	if err != nil {
		log.Println("Error starting rpc service:", err)
		return err
	}
	defer listen.Close()

	for {
		rpcConn, err := listen.Accept()
		if err != nil {
			log.Println("Error accepting rpc connection:", err)
			continue
		}
		go rpc.ServeConn(rpcConn)
	}
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

	log.Println("Connected to mongodb")

	return client, nil
}

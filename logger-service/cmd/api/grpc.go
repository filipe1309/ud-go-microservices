package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/filipe1309/ud-go-microservices/logger-service/data"
	"github.com/filipe1309/ud-go-microservices/logger-service/logs"
	"google.golang.org/grpc"
)

type LogServer struct {
	logs.UnimplementedLoggerServiceServer
	Models data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {
	input := req.GetLogEntry()
	
	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Models.LogEntry.Insert(logEntry)
	if err != nil {
		res := &logs.LogResponse{Result: "failed"}
		return res, err
	}

	res := &logs.LogResponse{Result: "logged"}
	return res, nil
}


func (app *Config) gRPCListen() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	server := grpc.NewServer()

	logs.RegisterLoggerServiceServer(server, &LogServer{Models: app.Models})

	log.Printf("gRPC server listening on port %s", gRpcPort)

	if err := server.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}

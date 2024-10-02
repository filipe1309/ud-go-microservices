package main

import (
	"context"

	"github.com/filipe1309/ud-go-microservices/logger-service/data"
	"github.com/filipe1309/ud-go-microservices/logger-service/logs"
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
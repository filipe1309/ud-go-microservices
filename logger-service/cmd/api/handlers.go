package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/filipe1309/ud-go-microservices/logger-service/data"
)

type JSONPayload struct {
	Name string      `json:"name"`
	Data   string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	var requestPayload JSONPayload
	_ = app.readJson(w, r, &requestPayload)

	// insert data
	event := data.LogEntry{
		Name: requestPayload.Name,
		Data: requestPayload.Data,
	}

	err := app.Models.LogEntry.Insert(event)
	if err != nil {
		app.errorJson(w, err)
	}

	res := jsonResponse{
		Error:  false,
		Message: "logged",
	}

	app.writeJson(w, http.StatusAccepted, res)
}


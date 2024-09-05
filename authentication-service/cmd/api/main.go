package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/filipe1309/ud-go-microservices/authentication-service/data"
)

const webPort = "80"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("Starting authentication service")

	// TODO: connect to DB

	// set up config
	app := Config{}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

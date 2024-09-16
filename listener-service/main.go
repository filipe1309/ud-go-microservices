package main

import (
	"fmt"
	"log"
	"math"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// try to connect to rabbitmq
	rabbitConn, err := connect()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	defer rabbitConn.Close()
	log.Println("Connected to RabbitMQ")

	// start listening for messages

	// create consumer

	// watch the queue for messages and consume events
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var conn *amqp.Connection

	// don't continue until rabbitmq is up
	for {
		c, err := amqp.Dial("amqp://guest:guest@localhost")
		if err != nil {
			fmt.Println("RabbitMQ not yet ready...")
			counts++
		} else {
			conn = c
			break
		}

		if counts > 5 {
			fmt.Println("RabbitMQ is not available. Exiting...", err)
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Printf("Retrying in %v seconds...", backOff)
		time.Sleep(backOff)
	}

	return conn, nil
}

package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/filipe1309/ud-go-microservices/listener-service/event"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// try to connect to rabbitmq
	rabbitConn, err := connect()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	defer rabbitConn.Close()

	// start listening for messages
	log.Println("Listening for messages...")

	// create consumer
	consumer, err := event.NewConsumer(rabbitConn)
	if err != nil {
		log.Panic(err)
	}

	// watch the queue for messages and consume events
	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})
	if err != nil {
		log.Println(err)
	}
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var conn *amqp.Connection

	// don't continue until rabbitmq is up
	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			fmt.Println("RabbitMQ not yet ready...")
			counts++
		} else {
			log.Println("Connected to RabbitMQ")
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

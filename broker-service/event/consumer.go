package event

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	ampq "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn      *ampq.Connection
	queueName string
}

func NewConsumer(conn *ampq.Connection) (Consumer, error) {
	consumer := Consumer{conn: conn}

	err := consumer.setup()
	if err != nil {
		return Consumer{}, err
	}

	return consumer, nil
}

func (consumer *Consumer) setup() error {
	// create a channel
	channel, err := consumer.conn.Channel()
	if err != nil {
		return err
	}

	return declareExchange(channel)
}

type Payload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (consumer *Consumer) Listen(topics []string) error {
	ch, err := consumer.conn.Channel()
	if err != nil {
		return err
	}

	defer ch.Close()

	queue, err := declareRandomQueue(ch)
	if err != nil {
		return err
	}

	for _, topicName := range topics {
		err := ch.QueueBind(
			queue.Name,   // name
			topicName,    // key
			"logs_topic", // exchange
			false,        // no-wait
			nil,          // arguments
		)
		if err != nil {
			return err
		}
	}

	messages, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		return err
	}

	forever := make(chan bool)
	go func() {
		for d := range messages {
			var paylod Payload
			_ = json.Unmarshal(d.Body, &paylod)
			log.Println("Received message: ", paylod)
			go handlePayload(paylod)
		}
	}()

	fmt.Printf(" [*] Waiting for message [Exchange, Queue] [%s, %s]\n", "logs_topic", queue.Name)
	<-forever

	return nil
}

func handlePayload(payload Payload) {
	switch payload.Name {
	case "log", "event":
		err := logEvent(payload)
		if err != nil {
			log.Println("Failed to log event: ", err)
		}
	// case "auth":
	default:
		err := logEvent(payload)
		if err != nil {
			log.Println("Failed to log event: ", err)
		}
	}
}

func logEvent(entry Payload) error {
	// create json that will be sent to logger service
	jsonData, _ := json.MarshalIndent(entry, "", "\t")

	// call logger service
	loggerServiceURL := "http://logger-service/log"
	request, err := http.NewRequest("POST", loggerServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(request)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// validate status code of response
	if res.StatusCode != http.StatusAccepted {
		return err
	}

	return nil
}

package event

import (
	"log"

	ampq "github.com/rabbitmq/amqp091-go"
)

type Emitter struct {
	conn *ampq.Connection
}

func NewEventEmitter(conn *ampq.Connection) (Emitter, error) {
	emitter := Emitter{conn: conn}

	err := emitter.setup()
	if err != nil {
		return Emitter{}, err
	}

	return emitter, nil
}

func (e *Emitter) setup() error {
	// create a channel
	channel, err := e.conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	return declareExchange(channel)
}

func (e *Emitter) Push(event, severity string) error {
	channel, err := e.conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	log.Println("Publishing to channel -> event", event, "with severity", severity)

	err = channel.Publish(
		"logs_topic", // exchange
		severity,     // routing key
		false,        // mandatory
		false,        // immediate
		ampq.Publishing{
			ContentType: "text/plain",
			Body:        []byte(event),
		},
	)

	return err
}

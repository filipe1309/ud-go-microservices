package event

import (
	ampq "github.com/rabbitmq/amqp091-go"
)

func declareExchange(channel *ampq.Channel) error {
	return channel.ExchangeDeclare(
		"logs_topic", // name
		"topic",      // type
		true,         // durable?
		false,        // auto-deleted?
		false,        // internal?
		false,        // no-wait?
		nil,          // arguments?
	)
}

func declareRandomQueue(channel *ampq.Channel) (ampq.Queue, error) {
	return channel.QueueDeclare(
		"",    // name?
		false, // durable?
		false, // delete when unused?
		true,  // exclusive?
		false, // no-wait?
		nil,   // arguments?
	)
}

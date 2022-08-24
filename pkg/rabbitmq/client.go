package rabbitmq

import "github.com/rabbitmq/amqp091-go"

type RabbitMQ struct {
	Channel *amqp091.Channel
}

func New(rabbitMqURL string) (Queue, error) {

	connection, err := amqp091.Dial(rabbitMqURL)
	if err != nil {
		return nil, err
	}

	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{Channel: channel}, nil
}

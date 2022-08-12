package rabbitmq

import "github.com/rabbitmq/amqp091-go"

type RabbitMQ struct {
	Channel *amqp091.Channel
}

func (mq *RabbitMQ) Connect(uri string) (*RabbitMQ, error) {

	conn, err := amqp091.Dial(uri)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{Channel: channel}, nil
}

func (mq *RabbitMQ) QueueDeclare(name string) (amqp091.Queue, error) {
	return mq.Channel.QueueDeclare(name, true, false, false, false, nil)
}

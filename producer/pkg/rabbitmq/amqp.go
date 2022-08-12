package rabbitmq

import "github.com/rabbitmq/amqp091-go"

type RabbitMQ struct {
	uri string
}

func New(uri string) *RabbitMQ {
	return &RabbitMQ{uri: uri}
}

func (mq *RabbitMQ) Connect() (*amqp091.Channel, error) {
	conn, err := amqp091.Dial(mq.uri)
	if err != nil {
		return nil, err
	}
	return conn.Channel()
}

func (mq *RabbitMQ) QueueDeclare(channel *amqp091.Channel, name string) (amqp091.Queue, error) {
	return channel.QueueDeclare(name, true, false, false, false, nil)
}

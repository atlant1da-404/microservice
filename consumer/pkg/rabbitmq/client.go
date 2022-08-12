package rabbitmq

import (
	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	uri         string
	uploadQueue string
}

func New(uri string) *RabbitMQ {
	return &RabbitMQ{uri: uri, uploadQueue: "upload"}
}

func (mq *RabbitMQ) Connect() (*amqp091.Channel, error) {
	conn, err := amqp091.Dial(mq.uri)
	if err != nil {
		return nil, err
	}
	return conn.Channel()
}

func (mq *RabbitMQ) QueueDeclare(channel *amqp091.Channel) (amqp091.Queue, error) {
	return channel.QueueDeclare(mq.uploadQueue, true, false, false, false, nil)
}

func MessageChan(channel *amqp091.Channel, name string) (<-chan amqp091.Delivery, error) {

	queue, err := channel.QueueDeclare(name, true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	if err := channel.Qos(1, 0, false); err != nil {
		return nil, err
	}

	return channel.Consume(queue.Name, "", false, false, false, false, nil)
}

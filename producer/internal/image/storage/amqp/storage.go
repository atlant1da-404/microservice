package amqp

import (
	"github.com/rabbitmq/amqp091-go"
	"producer/pkg/rabbitmq"
)

type rabbitMQ struct {
	client *rabbitmq.RabbitMQ
}

func NewRabbitMQ(uri string) (*rabbitMQ, error) {

	rb := rabbitmq.RabbitMQ{}

	client, err := rb.Connect(uri)
	if err != nil {
		return nil, err
	}

	return &rabbitMQ{client: client}, nil
}

func (mq *rabbitMQ) UploadImage(bFile []byte) error {

	queue, err := mq.client.QueueDeclare("upload")
	if err != nil {
		return err
	}

	return mq.client.Channel.Publish("", queue.Name, false, false, amqp091.Publishing{
		DeliveryMode: amqp091.Persistent,
		ContentType:  "application/json",
		Body:         bFile,
	})
}

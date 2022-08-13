package rabbitmq

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
)

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

func (mq *RabbitMQ) PublishWithContext(ctx context.Context, queueName string, contentType string, bFile []byte) error {
	return mq.Channel.PublishWithContext(ctx, "", queueName, false, false, amqp091.Publishing{
		DeliveryMode: amqp091.Persistent,
		ContentType:  contentType,
		Body:         bFile,
	})
}

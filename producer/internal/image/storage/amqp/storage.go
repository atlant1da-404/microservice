package amqp

import (
	"context"
	"producer/pkg/rabbitmq"
	"time"
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

func (mq *rabbitMQ) UploadImage(ctx context.Context, bFile []byte) error {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return mq.client.PublishWithContext(ctx, "upload", "application/json", bFile)
}

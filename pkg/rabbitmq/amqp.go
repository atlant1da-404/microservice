package rabbitmq

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
)

// Queue interface amqp transport level of service
type Queue interface {
	// Send into queue
	Send(ctx context.Context, bFile []byte, queueName string, contentType string) error
	// Listen from queue
	Listen(name string) (<-chan amqp091.Delivery, error)
}

func (mq *RabbitMQ) queueDeclare(channel *amqp091.Channel, queueName string) (amqp091.Queue, error) {
	return channel.QueueDeclare(queueName, true, false, false, false, nil)
}

func (mq *RabbitMQ) Send(ctx context.Context, bFile []byte, queueName string, contentType string) error {

	queue, err := mq.queueDeclare(mq.Channel, queueName)
	if err != nil {
		return err
	}

	return mq.Channel.PublishWithContext(ctx, "", queue.Name, false, false, amqp091.Publishing{
		DeliveryMode: amqp091.Persistent,
		ContentType:  contentType,
		Body:         bFile,
	})
}

func (mq *RabbitMQ) Listen(name string) (<-chan amqp091.Delivery, error) {

	queue, err := mq.Channel.QueueDeclare(name, true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	if err := mq.Channel.Qos(1, 0, false); err != nil {
		return nil, err
	}

	return mq.Channel.Consume(queue.Name, "", false, false, false, false, nil)
}

package rabbitmq

import (
	"context"
	"github.com/rabbitmq/amqp091-go"
)

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

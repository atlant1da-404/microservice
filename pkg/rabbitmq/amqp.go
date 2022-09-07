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

package amqp

import (
	"log"
	"producer/pkg/rabbitmq"

	"github.com/rabbitmq/amqp091-go"
)

type rabbitMQ struct {
	rabbitPkg *rabbitmq.RabbitMQ
}

func NewRabbitMQ(uri string) *rabbitMQ {
	amqp := rabbitmq.New(uri)
	return &rabbitMQ{rabbitPkg: amqp}
}

func (mq *rabbitMQ) UploadImage(bData []byte) error {

	channel, err := mq.rabbitPkg.Connect()
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer channel.Close()

	queue, err := mq.rabbitPkg.QueueDeclare(channel)
	if err != nil {
		return err
	}

	return channel.Publish("", queue.Name, false, false, amqp091.Publishing{
		DeliveryMode: amqp091.Persistent,
		ContentType:  "application/json",
		Body:         bData,
	})
}

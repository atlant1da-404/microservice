package amqp

import (
	"github.com/rabbitmq/amqp091-go"
	"log"
	"producer/pkg/rabbitmq"
)

type rabbitMQ struct {
	rabbitPkg *rabbitmq.RabbitMQ
	channel   *amqp091.Channel
}

func NewRabbitMQ(uri string) *rabbitMQ {
	amqp := rabbitmq.New(uri)
	channel, err := amqp.Connect()
	if err != nil {
		log.Fatalln(err.Error())
	}
	return &rabbitMQ{rabbitPkg: amqp, channel: channel}
}

func (mq *rabbitMQ) UploadImage(bData []byte) error {

	queue, err := mq.rabbitPkg.QueueDeclare(mq.channel, "upload")
	if err != nil {
		return err
	}

	return mq.channel.Publish("", queue.Name, false, false, amqp091.Publishing{
		DeliveryMode: amqp091.Persistent,
		ContentType:  "application/json",
		Body:         bData,
	})
}

func (mq *rabbitMQ) DownloadImage(bData []byte) ([]byte, error) {

	return nil, nil
}

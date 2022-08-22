package rabbitmq

type Queue interface {
	Send()
	Listen()
}

func (mq *RabbitMQ) Send() {
	return
}

func (mq *RabbitMQ) Listen() {
	return
}

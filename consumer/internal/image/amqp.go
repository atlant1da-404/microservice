package image

import (
	"consumer/pkg/rabbitmq"
	"github.com/rabbitmq/amqp091-go"
	"log"
	"sync"
)

type Handler struct {
	Channel      *amqp091.Channel
	ImageService Service
}

func (h *Handler) Consumer(wg *sync.WaitGroup) {
	defer wg.Done()

	messageCh, err := rabbitmq.MessageChan(h.Channel)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	wg.Add(1)
	go func(wg *sync.WaitGroup, messageChan <-chan amqp091.Delivery) {
		defer wg.Done()

		for message := range messageChan {

			if err := h.ImageService.OptimizeImage(message.Body); err != nil {
				log.Printf("error photo processing: %s", err.Error())
			}

			if err := message.Ack(false); err != nil {
				log.Printf("error acknowledging message : %s", err)
			} else {
				log.Printf("acknowledged message")
			}
		}
	}(wg, messageCh)
	wg.Wait()
}

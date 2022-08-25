package img

import (
	"github.com/rabbitmq/amqp091-go"
	"images/internal/service"
	"images/pkg/rabbitmq"
	"log"
	"sync"
)

type ImageAmqpHandler struct {
	Amqp         rabbitmq.Queue
	ImageService service.ImageService
}

func NewImageAmqpHandler(amqp rabbitmq.Queue, imageService service.ImageService) *ImageAmqpHandler {
	return &ImageAmqpHandler{
		Amqp:         amqp,
		ImageService: imageService,
	}
}

func (a *ImageAmqpHandler) Register() {
	go a.Listen()
}

// Listen amqp queue
// Queue name: "upload"
func (a *ImageAmqpHandler) Listen() {

	messageCh, err := a.Amqp.Listen("upload")
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup, messageCh <-chan amqp091.Delivery) {
		defer wg.Done()

		for message := range messageCh {

			log.Print("[CONSUMER]: Accept upload image")

			if err := a.ImageService.SaveImage(message.Body); err != nil {
				log.Printf("[CONSUMER]: error saving image: %s", err.Error())
			}

			if err := message.Ack(false); err != nil {
				log.Printf("[CONSUMER]: error acknowledging message : %s", err.Error())
			} else {
				log.Printf("[CONSUMER]: Success")
			}
		}
	}(wg, messageCh)
	wg.Wait()
}

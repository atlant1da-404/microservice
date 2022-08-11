package main

import (
	"consumer/internal/config"
	"consumer/internal/image"
	"consumer/pkg/rabbitmq"
	"log"
	"sync"
)

func main() {

	cfg := config.GetConfig()
	wg := &sync.WaitGroup{}

	rabbitMq := rabbitmq.New(cfg.RabbitMQ)
	channel, err := rabbitMq.Connect()

	defer channel.Close()
	if err != nil {
		log.Fatalln(err.Error())
	}

	imageService := image.NewService()
	imageHandler := image.Handler{Channel: channel, ImageService: imageService}
	wg.Add(1)
	go imageHandler.Consumer(wg)
	log.Println("successfully started!")
	wg.Wait()
}

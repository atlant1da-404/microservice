package main

import (
	"consumer/internal/config"
	"consumer/internal/image"
	"consumer/internal/image/minio"
	"consumer/pkg/rabbitmq"
	"log"
	"sync"
)

func main() {

	cfg := config.GetConfig()
	wg := &sync.WaitGroup{}

	rabbitMq := rabbitmq.New(cfg.RabbitMQ)
	channel, err := rabbitMq.Connect()
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer channel.Close()

	imageStorage, err := minio.NewStorage(cfg.Minio, cfg.MinioAccessKey, cfg.MinioPassword)
	if err != nil {
		log.Fatalln(err.Error())
	}

	imageService := image.NewService(imageStorage)
	imageHandler := image.Handler{Channel: channel, ImageService: imageService}

	wg.Add(1)
	go imageHandler.Consumer(wg)
	log.Println("[CONSUMER]: successfully started!")
	wg.Wait()
}

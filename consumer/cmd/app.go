package main

import (
	"consumer/internal/config"
	"consumer/internal/image"
	"consumer/internal/image/storage/minio"
	"consumer/pkg/rabbitmq"
	"log"
)

func main() {

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalln(err.Error())
	}

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
	imageHandler.Consumer()
}

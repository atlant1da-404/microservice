package main

import (
	"consumer/internal/config"
	"consumer/internal/image"
	"consumer/internal/image/cache"
	"consumer/pkg/rabbitmq"
	"consumer/pkg/redis"
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

	rdb, err := redis.NewRedis(cfg.Redis, cfg.RedisPassword)
	if err != nil {
		log.Fatalln(err.Error())
	}

	imageStorage := cache.NewStorage(rdb)
	imageService := image.NewService(imageStorage)
	imageHandler := image.Handler{Channel: channel, ImageService: imageService}
	wg.Add(1)
	go imageHandler.Consumer(wg)
	log.Println("successfully started!")
	wg.Wait()
}

package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"producer/internal/config"
	"producer/internal/image"
	"producer/internal/image/storage/amqp"
	"producer/internal/image/storage/minio"
	"time"
)

func main() {

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalln(err.Error())
	}

	amqpStorage, err := amqp.NewRabbitMQ(cfg.RabbitMQ)
	if err != nil {
		log.Fatalln(err.Error())
	}

	minioStorage, err := minio.NewStorage(cfg.Minio, cfg.MinioAccessKey, cfg.MinioPassword)
	if err != nil {
		log.Fatalln(err.Error())
	}

	router := httprouter.New()
	imageService := image.NewService(amqpStorage, minioStorage)
	imageHandler := image.Handler{ImageService: imageService}
	imageHandler.Register(router)

	start(router, cfg)
}

func start(router http.Handler, cfg *config.Config) {

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
	if err != nil {
		log.Fatalln(err.Error())
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("[PRODUCER]: started on %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)

	if err := server.Serve(listener); err != nil {
		log.Fatalln(err.Error())
	}
}

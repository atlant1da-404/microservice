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

	cfg := config.GetConfig()
	router := httprouter.New()
	errCh := make(chan error)

	amqpStorage, err := amqp.NewRabbitMQ(cfg.RabbitMQ)
	if err != nil {
		log.Fatalln(err.Error())
	}

	minioStorage, err := minio.NewStorage(cfg.Minio, cfg.MinioAccessKey, cfg.MinioPassword)
	if err != nil {
		log.Fatalln(err.Error())
	}

	imageService := image.NewService(amqpStorage, minioStorage)
	imageHandler := image.Handler{ImageService: imageService}
	imageHandler.Register(router)

	go start(router, cfg, errCh)

	log.Println("[PRODUCER] successfully started on localhost:8081")
	log.Fatalln(<-errCh)
}

func start(router http.Handler, cfg *config.Config, errCh chan error) {

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
	if err != nil {
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	errCh <- server.Serve(listener)
}

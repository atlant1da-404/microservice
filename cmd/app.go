package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"images/internal/config"
	"images/internal/service/image"
	"images/internal/storage/image/minio"
	"images/internal/transport/amqp/image"
	"images/internal/transport/http/controllers/v1"
	"images/pkg/minio"
	"images/pkg/rabbitmq"
	"images/pkg/resize"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {

	cfg, err := config.GetConfig("dev")
	if err != nil {
		log.Fatalf("[FAIL]: error reading config: %s", err.Error())
	}

	if err := start(cfg); err != nil {
		log.Fatalf("[FAIL]: server error %s", err.Error())
	}

}

func start(cfg *config.Config) error {

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.BindIP, cfg.Port))
	if err != nil {
		return err
	}

	amqpPkg, err := rabbitmq.New(cfg.RabbitMQ)
	if err != nil {
		return err
	}

	minioPkg, err := minio.NewMinio(cfg.Minio, cfg.MinioAccessKey, cfg.MinioPassword)
	if err != nil {
		return err
	}

	router := httprouter.New()
	resizer := resize.NewCompressor()

	imageStorage := storage.NewImageStorage(minioPkg)
	imageService := image.NewImageService(imageStorage, amqpPkg, resizer)

	imageAmqp := img.NewImageAmqpHandler(amqpPkg, imageService)
	imageHandler := v1.NewImageHandler(imageService)

	imageHandler.Register(router)
	imageAmqp.Register()

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("[SERVER] Started at %s:%s", cfg.BindIP, cfg.Port)
	return server.Serve(listener)
}

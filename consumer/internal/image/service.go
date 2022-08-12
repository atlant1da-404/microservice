package image

import (
	"consumer/pkg/convert"
	"consumer/pkg/resize"
	"encoding/json"
	"fmt"
	"github.com/rabbitmq/amqp091-go"
	"os"
)

type service struct {
	ImageStorage Storage
}

func NewService(storage Storage) Service {
	return &service{
		ImageStorage: storage,
	}
}

type Service interface {
	OptimizeImage(data []byte) error
	Base64Enc(data []byte) (*SendDTO, error)
	SendImage(channel *amqp091.Channel, sendDto *SendDTO) error
}

func (s *service) OptimizeImage(data []byte) error {

	model := &UploadDTO{}
	if err := json.Unmarshal(data, &model); err != nil {
		return err
	}

	img, err := convert.Base64Dec(model.Base64)
	if err != nil {
		return err
	}

	if err := resize.Resize(img, model.Id); err != nil {
		return err
	}

	return s.ImageStorage.Set(model.Id)
}

func (s *service) Base64Enc(data []byte) (*SendDTO, error) {

	model := &DownloadDTO{}
	if err := json.Unmarshal(data, model); err != nil {
		return nil, err
	}

	filename := fmt.Sprintf("%d_%s.jpg", model.ID, model.Quality)
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	base64Enc, err := convert.Base64Enc(file)
	if err != nil {
		return nil, err
	}

	return &SendDTO{Id: model.ID, Base64: base64Enc}, nil
}

func (s *service) SendImage(channel *amqp091.Channel, sendDto *SendDTO) error {

	bData, err := json.Marshal(sendDto)
	if err != nil {
		return err
	}

	queue, err := channel.QueueDeclare("producer", true, false, false, false, nil)
	if err != nil {
		return err
	}

	return channel.Publish("", queue.Name, false, false, amqp091.Publishing{
		DeliveryMode: amqp091.Persistent,
		ContentType:  "application/json",
		Body:         bData,
	})
}

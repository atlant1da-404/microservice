package service

import (
	"context"
	"images/internal/model"
)

// ImageService is realization of processing image for business logic
type ImageService interface {
	// CreateImage performs file to byte and sends to amqp
	CreateImage(ctx context.Context, dto model.UploadFileDTO) error
	// SaveImage compress the image and send to storage
	SaveImage(data []byte) error
	// GetImage a validation for request from client and returns a image from storage in *model.File format
	GetImage(ctx context.Context, dto model.DownloadFileDTO) (*model.File, error)
}

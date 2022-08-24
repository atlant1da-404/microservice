package service

import (
	"context"
	"images/internal/model"
)

type ImageService interface {
	CreateImage(ctx context.Context, dto model.UploadFileDTO) error
	SaveImage(data []byte) error
	GetImage(ctx context.Context, dto model.DownloadFileDTO) (*model.File, error)
}

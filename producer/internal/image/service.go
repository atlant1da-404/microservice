package image

import (
	"context"
	"encoding/json"
	"fmt"
	"producer/internal/apperror"
)

type service struct {
	StorageAmqp  StorageAmqp
	StorageMinio StorageMinio
}

func NewService(storageAmqp StorageAmqp, storageMinio StorageMinio) Service {
	return &service{StorageAmqp: storageAmqp, StorageMinio: storageMinio}
}

type Service interface {
	UploadImage(ctx context.Context, dto UploadFileDTO) error
	DownloadImage(ctx context.Context, dto DownloadFileDTO) (*File, error)
}

func (s *service) UploadImage(ctx context.Context, dto UploadFileDTO) error {

	img, err := NewFile(dto)
	if err != nil {
		return apperror.BadRequestError(err.Error())
	}

	bFile, err := json.Marshal(img)
	if err != nil {
		return apperror.BadRequestError(err.Error())
	}

	return s.StorageAmqp.UploadImage(ctx, bFile)
}

func (s *service) DownloadImage(ctx context.Context, dto DownloadFileDTO) (*File, error) {

	if err := dto.Validate(); err != nil {
		return nil, apperror.BadRequestError(err.Error())
	}

	return s.StorageMinio.DownloadImage(ctx, fmt.Sprintf("%s_%s.jpeg", dto.ID, dto.Quality))
}

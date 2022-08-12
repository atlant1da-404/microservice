package image

import (
	"encoding/json"
	"fmt"
)

type service struct {
	StorageAmqp  StorageAmqp
	StorageMinio StorageMinio
}

func NewService(storageAmqp StorageAmqp, storageMinio StorageMinio) Service {
	return &service{StorageAmqp: storageAmqp, StorageMinio: storageMinio}
}

type Service interface {
	UploadImage(dto UploadFileDTO) error
	DownloadImage(dto DownloadFileDTO) (*File, error)
}

func (s *service) UploadImage(dto UploadFileDTO) error {

	img, err := NewFile(dto)
	if err != nil {
		return err
	}

	bFile, err := json.Marshal(img)
	if err != nil {
		return err
	}

	return s.StorageAmqp.UploadImage(bFile)
}

func (s *service) DownloadImage(dto DownloadFileDTO) (*File, error) {

	if err := dto.Validate(); err != nil {
		return nil, err
	}

	return s.StorageMinio.DownloadImage(fmt.Sprintf("%s_%s.jpeg", dto.ID, dto.Quality))
}

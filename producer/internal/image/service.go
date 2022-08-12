package image

import (
	"encoding/json"
)

type service struct {
	StorageAmqp StorageAmqp
}

func NewService(storageAmqp StorageAmqp) Service {
	return &service{StorageAmqp: storageAmqp}
}

type Service interface {
	UploadImage(dto UploadFileDTO) error
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

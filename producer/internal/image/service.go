package image

import (
	"encoding/json"
	"producer/pkg/convert"
	"time"
)

type service struct {
	ImageStorage Storage
}

func NewService(storage Storage) Service {
	return &service{ImageStorage: storage}
}

type Service interface {
	UploadImage(dto UploadDTO) error
	DownloadImage(uuid int) error
}

func (s *service) UploadImage(dto UploadDTO) error {

	if err := dto.Validate(); err != nil {
		return err
	}

	base64Enc, err := convert.Base64Enc(dto.File)
	if err != nil {
		return err
	}

	dto.Id = time.Now().UTC().UnixMilli()
	dto.Base64 += base64Enc
	dto.File = nil

	bData, err := json.Marshal(&dto)
	if err != nil {
		return err
	}

	return s.ImageStorage.UploadImage(bData)
}

func (s *service) DownloadImage(uuid int) error {
	return nil
}

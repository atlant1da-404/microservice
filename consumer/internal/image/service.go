package image

import (
	"consumer/pkg/convert"
	"consumer/pkg/resize"
	"encoding/json"
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

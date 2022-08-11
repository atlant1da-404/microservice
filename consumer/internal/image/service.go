package image

import (
	"consumer/pkg/convert"
	"consumer/pkg/resize"
	"encoding/json"
)

type service struct {
}

func NewService() Service {
	return &service{}
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

	resize.Resize(img, model.Id)
	return nil
}

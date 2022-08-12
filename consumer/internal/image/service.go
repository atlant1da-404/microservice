package image

import (
	"bytes"
	"consumer/pkg/resize"
	"encoding/json"
	"image/jpeg"
)

type service struct {
	imageStorage Storage
}

func NewService(imageStorage Storage) Service {
	return &service{imageStorage: imageStorage}
}

type Service interface {
	SaveImage(data []byte) error
}

func (s *service) SaveImage(data []byte) error {

	var file File

	if err := json.Unmarshal(data, &file); err != nil {
		return err
	}

	img, err := resize.GetImage(file.Bytes)
	if err != nil {
		return err
	}

	pictures, quality := resize.ImageQuality(img)

	for i, picture := range pictures {

		fileID := resize.GeneratePictureId(file.ID, quality[i])

		buf := new(bytes.Buffer)
		if err := jpeg.Encode(buf, picture, nil); err != nil {
			return err
		}

		reader := bytes.NewReader(buf.Bytes())
		if err := s.imageStorage.UploadImage(fileID, reader.Size(), reader); err != nil {
			return err
		}
	}

	return err
}

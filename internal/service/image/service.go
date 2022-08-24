package image

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"image/jpeg"
	"images/internal/model"
	"images/internal/service"
	"images/internal/storage"
	"images/pkg/apperror"
	"images/pkg/rabbitmq"
	"images/pkg/resize"
)

type imageService struct {
	ImageStorage storage.ImageStorage
	Amqp         rabbitmq.Queue
	Resize       resize.Compressor
}

func NewImageService(imageStorage storage.ImageStorage, amqp rabbitmq.Queue, resizer resize.Compressor) service.ImageService {
	return &imageService{
		ImageStorage: imageStorage,
		Amqp:         amqp,
		Resize:       resizer,
	}
}

func (s *imageService) CreateImage(ctx context.Context, dto model.UploadFileDTO) error {

	file, err := dto.NewFile()
	if err != nil {
		return apperror.SystemError(err.Error())
	}

	bFile, err := json.Marshal(file)
	if err != nil {
		return apperror.BadRequestError(err.Error())
	}

	return s.Amqp.Send(ctx, bFile, "upload", "application/json")
}

func (s *imageService) SaveImage(data []byte) error {

	var file model.File

	if err := json.Unmarshal(data, &file); err != nil {
		return err
	}

	image, err := s.Resize.GetImage(file.Bytes)
	if err != nil {
		return err
	}

	pictures, quality := s.Resize.Compress(image)

	for i, picture := range pictures {

		fileID := s.Resize.GeneratePictureId(file.ID, quality[i])

		buf := new(bytes.Buffer)
		if err := jpeg.Encode(buf, picture, nil); err != nil {
			return err
		}

		reader := bytes.NewReader(buf.Bytes())
		if err := s.ImageStorage.SaveFile(fileID, reader.Size(), reader); err != nil {
			return err
		}
	}

	return err
}

func (s *imageService) GetImage(ctx context.Context, dto model.DownloadFileDTO) (*model.File, error) {

	if err := dto.Validate(); err != nil {
		return nil, apperror.BadRequestError(err.Error())
	}

	return s.ImageStorage.GetFile(ctx, fmt.Sprintf("%s_%s.jpeg", dto.Id, dto.Quality))
}

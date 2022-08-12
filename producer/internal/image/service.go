package image

import (
	"encoding/json"
	"producer/internal/apperror"
	"producer/pkg/convert"
	"time"
)

type service struct {
	StorageAmqp  StorageAmqp
	StorageCache StorageCache
}

func NewService(storageAmqp StorageAmqp, storageCache StorageCache) Service {
	return &service{StorageAmqp: storageAmqp, StorageCache: storageCache}
}

type Service interface {
	UploadImage(dto UploadDTO) (int64, error)
	DownloadImage(dto DownloadDTO) (*DownloadImage, error)
}

func (s *service) UploadImage(dto UploadDTO) (int64, error) {

	if err := dto.Validate(); err != nil {
		return 0, err
	}

	base64Enc, err := convert.Base64Enc(dto.File)
	if err != nil {
		return 0, err
	}

	dto.Id = time.Now().UTC().UnixMilli()
	dto.Base64 += base64Enc
	dto.File = nil

	bData, err := json.Marshal(&dto)
	if err != nil {
		return 0, err
	}

	if err := s.StorageAmqp.UploadImage(bData); err != nil {
		return 0, err
	}

	return dto.Id, nil
}

func (s *service) DownloadImage(dto DownloadDTO) (*DownloadImage, error) {

	if err := dto.Validate(); err != nil {
		return nil, err
	}

	if notExist := s.StorageCache.CheckInCache(dto.Id); !notExist {
		return nil, apperror.ErrNotFound
	}

	bData, err := json.Marshal(&dto)
	if err != nil {
		return nil, err
	}

	data, err := s.StorageAmqp.DownloadImage(bData)
	if err != nil {
		return nil, err
	}

	var model = &DownloadImage{}
	if err := json.Unmarshal(data, model); err != nil {
		return nil, err
	}

	return model, nil
}

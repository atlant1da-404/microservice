package storage

import (
	"context"
	"images/internal/model"
	"images/internal/storage"
	"images/pkg/minio"
	"io"
)

type imageStorage struct {
	Minio minio.Minio
}

func NewImageStorage(minio minio.Minio) storage.ImageStorage {
	return &imageStorage{Minio: minio}
}

func (s *imageStorage) SaveFile(fileId string, size int64, reader io.Reader) error {
	return s.Minio.Save(context.Background(), "upload", fileId, size, reader)
}

func (s *imageStorage) GetFile(ctx context.Context, fileId string) (*model.File, error) {

	obj, err := s.Minio.Get(ctx, "upload", fileId)
	if err != nil {
		return nil, err
	}
	defer obj.Close()

	objectInfo, err := obj.Stat()
	if err != nil {
		return nil, err
	}

	buffer := make([]byte, objectInfo.Size)
	_, err = obj.Read(buffer)
	if err != nil && err != io.EOF {
		return nil, err
	}

	return &model.File{ID: objectInfo.Key, Size: objectInfo.Size, Bytes: buffer}, nil
}

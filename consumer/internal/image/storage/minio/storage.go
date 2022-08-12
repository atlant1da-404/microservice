package minio

import (
	"consumer/internal/image"
	"consumer/pkg/minio"
	"fmt"
	"io"
)

type minioStorage struct {
	client *minio.Client
}

func NewStorage(endpoint, accessKeyID, secretAccessKey string) (image.Storage, error) {

	client, err := minio.NewClient(endpoint, accessKeyID, secretAccessKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create minio client. err: %w", err)
	}

	return &minioStorage{
		client: client,
	}, nil
}

func (m *minioStorage) UploadImage(fileId string, fileSize int64, reader io.Reader) error {
	return m.client.UploadImage(fileId, fileSize, reader)
}

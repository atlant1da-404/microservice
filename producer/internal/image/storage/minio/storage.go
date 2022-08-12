package minio

import (
	"fmt"
	"io"
	"producer/internal/image"
	"producer/pkg/minio"
)

type minioStorage struct {
	client *minio.Client
}

func NewStorage(endpoint, accessKeyID, secretAccessKey string) (image.StorageMinio, error) {

	client, err := minio.NewClient(endpoint, accessKeyID, secretAccessKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create minio client. err: %w", err)
	}

	return &minioStorage{
		client: client,
	}, nil
}

func (m *minioStorage) DownloadImage(fileId string) (*image.File, error) {

	obj, err := m.client.Download("upload", fileId)
	if err != nil {
		return nil, fmt.Errorf("failed to get file. err: %w", err)
	}
	defer obj.Close()

	objectInfo, err := obj.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to get file. err: %w", err)
	}

	buffer := make([]byte, objectInfo.Size)
	_, err = obj.Read(buffer)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("failed to get objects. err: %w", err)
	}

	f := image.File{
		ID:    objectInfo.Key,
		Size:  objectInfo.Size,
		Bytes: buffer,
	}

	return &f, nil
}

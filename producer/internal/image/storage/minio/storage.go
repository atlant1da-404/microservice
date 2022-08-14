package minio

import (
	"context"
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

	return &minioStorage{client: client}, nil
}

func (m *minioStorage) DownloadImage(ctx context.Context, fileId string) (*image.File, error) {

	obj, err := m.client.Download(ctx, "upload", fileId)
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

	return &image.File{ID: objectInfo.Key, Size: objectInfo.Size, Bytes: buffer}, nil
}

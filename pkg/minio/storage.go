package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"io"
	"time"
)

// Minio package interface
type Minio interface {
	// Save with context a data in bucket
	Save(ctx context.Context, bucketName string, fileId string, size int64, reader io.Reader) error
	// Get with context a data from bucket
	Get(ctx context.Context, bucketName string, fileID string) (*minio.Object, error)
}

func (m *Storage) Save(ctx context.Context, bucketName, fileId string, size int64, reader io.Reader) error {

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	exists, err := m.minioClient.BucketExists(ctx, bucketName)
	if err != nil || !exists {
		err := m.minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return fmt.Errorf("failed to create new bucket. err: %w", err)
		}
	}

	opt := minio.PutObjectOptions{ContentType: "application/octet-stream"}
	_, errs := m.minioClient.PutObject(ctx, bucketName, fileId, reader, size, opt)
	return errs
}

func (m *Storage) Get(ctx context.Context, bucketName string, fileID string) (*minio.Object, error) {
	return m.minioClient.GetObject(ctx, bucketName, fileID, minio.GetObjectOptions{})
}

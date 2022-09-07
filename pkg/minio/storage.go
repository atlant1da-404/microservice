package minio

import (
	"context"
	"io"

	"github.com/minio/minio-go/v7"
)

// Minio package interface
type Minio interface {
	// Save with context a data in bucket
	Save(ctx context.Context, bucketName string, fileId string, size int64, reader io.Reader) error
	// Get with context a data from bucket
	Get(ctx context.Context, bucketName string, fileID string) (*minio.Object, error)
}

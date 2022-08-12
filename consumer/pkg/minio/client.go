package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"time"
)

const bucketName = "upload"

type Client struct {
	minioClient *minio.Client
}

func NewClient(endpoint, accessKeyID, secretAccessKey string) (*Client, error) {

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create minio client. err: %w", err)
	}

	return &Client{minioClient: minioClient}, nil
}

func (c *Client) UploadImage(fileId string, fileSize int64, reader io.Reader) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	exists, err := c.minioClient.BucketExists(ctx, bucketName)
	if err != nil || !exists {
		err := c.minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return fmt.Errorf("failed to create new bucket. err: %w", err)
		}
	}

	opt := minio.PutObjectOptions{ContentType: "application/octet-stream"}
	_, errs := c.minioClient.PutObject(ctx, bucketName, fileId, reader, fileSize, opt)
	return errs
}

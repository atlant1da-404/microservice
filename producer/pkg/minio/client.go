package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

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

func (c *Client) Download(ctx context.Context, bucketName, fileId string) (*minio.Object, error) {
	return c.minioClient.GetObject(ctx, bucketName, fileId, minio.GetObjectOptions{})
}

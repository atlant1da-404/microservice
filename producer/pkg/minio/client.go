package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
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

func (c *Client) Download(bucketName, fileId string) (*minio.Object, error) {

	obj, err := c.minioClient.GetObject(context.Background(), bucketName, fileId, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get file with id: %s from minio bucket %s. err: %w", fileId, bucketName, err)
	}
	return obj, nil
}

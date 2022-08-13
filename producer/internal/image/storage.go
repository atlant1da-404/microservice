package image

import "context"

type StorageAmqp interface {
	UploadImage(ctx context.Context, bFile []byte) error
}

type StorageMinio interface {
	DownloadImage(ctx context.Context, fileId string) (*File, error)
}

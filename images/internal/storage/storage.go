package storage

import (
	"context"
	"images/internal/model"
	"io"
)

type ImageStorage interface {
	SaveFile(fileId string, size int64, reader io.Reader) error
	GetFile(ctx context.Context, fileId string) (*model.File, error)
}

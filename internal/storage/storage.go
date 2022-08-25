package storage

import (
	"context"
	"images/internal/model"
	"io"
)

// ImageStorage is interface for image storage
type ImageStorage interface {
	// SaveFile save file to storage
	SaveFile(fileId string, size int64, reader io.Reader) error
	// GetFile from storage
	GetFile(ctx context.Context, fileId string) (*model.File, error)
}

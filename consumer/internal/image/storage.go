package image

import "io"

type Storage interface {
	UploadImage(fileId string, fileSize int64, reader io.Reader) error
}

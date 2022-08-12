package image

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"strconv"
	"time"
)

type File struct {
	ID    string `json:"id"`
	Size  int64  `json:"size"`
	Bytes []byte `json:"file"`
}

type UploadFileDTO struct {
	Id     int64          `json:"id"`
	Size   int64          `json:"size"`
	Reader multipart.File `json:"reader"`
}

func NewFile(dto UploadFileDTO) (*File, error) {

	bytes, err := ioutil.ReadAll(dto.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to create file model. err: %w", err)
	}

	return &File{
		ID:    strconv.FormatInt(time.Now().UTC().UnixMilli(), 10),
		Size:  dto.Size,
		Bytes: bytes,
	}, nil
}

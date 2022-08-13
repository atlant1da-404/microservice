package image

import (
	"errors"
	"io/ioutil"
	"mime/multipart"
	"producer/internal/apperror"
	"strconv"
)

type File struct {
	ID    string `json:"id"`
	Size  int64  `json:"size"`
	Bytes []byte `json:"file"`
}

type UploadFileDTO struct {
	Id     int            `json:"id"`
	Size   int64          `json:"size"`
	Reader multipart.File `json:"reader"`
}

func NewFile(dto UploadFileDTO) (*File, error) {

	bytes, err := ioutil.ReadAll(dto.Reader)
	if err != nil {
		return nil, apperror.SystemError(err.Error())
	}

	return &File{
		ID:    strconv.Itoa(dto.Id),
		Size:  dto.Size,
		Bytes: bytes,
	}, nil
}

type DownloadFileDTO struct {
	ID      string `json:"id"`
	Quality string `json:"quality"`
}

func (d *DownloadFileDTO) Validate() error {

	if d.ID == "" {
		return errors.New("id not found")
	}
	if d.Quality == "" {
		return errors.New("quality not found")
	}
	return nil
}

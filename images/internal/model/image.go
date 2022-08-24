package model

import (
	"errors"
	"images/pkg/apperror"
	"io/ioutil"
	"mime/multipart"
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

type DownloadFileDTO struct {
	Id      string `json:"id"`
	Quality string `json:"quality"`
}

func (d *DownloadFileDTO) Validate() error {

	if d.Id == "" {
		return errors.New("id not found")
	}

	for _, quality := range []string{"25", "50", "75", "100"} {
		if d.Quality == quality {
			return nil
		}
	}

	return errors.New("quality not correct")
}

func (u *UploadFileDTO) NewFile() (*File, error) {

	bytes, err := ioutil.ReadAll(u.Reader)
	if err != nil {
		return nil, apperror.SystemError(err.Error())
	}

	return &File{
		ID:    strconv.Itoa(u.Id),
		Size:  u.Size,
		Bytes: bytes,
	}, nil
}

package image

import (
	"mime/multipart"
	"producer/internal/apperror"
)

type UploadDTO struct {
	Id         int64                 `json:"id"`
	File       multipart.File        `json:"file"`
	FileHeader *multipart.FileHeader `json:"file_header"`
	Base64     string                `json:"base_64"`
}

func (u UploadDTO) Validate() error {
	if u.File == nil {
		return apperror.BadRequestError("image not found")
	}

	if u.FileHeader == nil {
		return apperror.BadRequestError("image not found")
	}

	return nil
}

type DownloadDTO struct {
	Id      int64  `json:"id"`
	Quality string `json:"quality"`
}

func (d DownloadDTO) Validate() error {
	if d.Id == 0 {
		return apperror.BadRequestError("id not correct")
	}
	if d.Quality == "" {
		return apperror.BadRequestError("quality not correct")
	}

	return nil
}

type DownloadImage struct {
	Base64 string `json:"base_64"`
}

package image

import "mime/multipart"

type UploadDTO struct {
	Id         int                   `json:"id"`
	FileHeader *multipart.FileHeader `json:"file_header"`
	Base64     string                `json:"base_64"`
}

type DownloadDTO struct {
	ID      int    `json:"id"`
	Quality string `json:"quality"`
}

type SendDTO struct {
	Id     int    `json:"id"`
	Base64 string `json:"base_64"`
}

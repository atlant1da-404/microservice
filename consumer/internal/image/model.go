package image

import "mime/multipart"

type UploadDTO struct {
	Id         int                   `json:"id"`
	FileHeader *multipart.FileHeader `json:"file_header"`
	Base64     string                `json:"base_64"`
}

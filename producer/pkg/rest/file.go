package rest

import (
	"errors"
	"mime/multipart"
	"net/http"
)

func GetFile(r *http.Request, formName string) (*multipart.FileHeader, error) {

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		return nil, err
	}

	files, ok := r.MultipartForm.File[formName]
	if !ok || len(files) == 0 {
		return nil, errors.New("file required")
	}

	return files[0], nil
}

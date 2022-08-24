package rest

import (
	"errors"
	"fmt"
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

func SendFile(w http.ResponseWriter, r *http.Request, fileId string, bFile []byte) error {

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileId))
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	_, err := w.Write(bFile)
	return err
}

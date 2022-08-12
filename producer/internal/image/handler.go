package image

import (
	"encoding/json"
	"net/http"
	"producer/internal/apperror"
	"producer/pkg/rest"
	"time"

	"github.com/julienschmidt/httprouter"
)

const (
	imageUploadURL   = "/api/image/upload"
	imageDownloadURL = "/api/image/download/:id"
)

type Handler struct {
	ImageService Service
}

func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, imageUploadURL, apperror.Middleware(h.UploadImage))
}

func (h *Handler) UploadImage(w http.ResponseWriter, r *http.Request) error {

	w.Header().Set("Content-Type", "form/json")
	file, err := rest.GetFile(r, "image")
	if err != nil {
		return apperror.BadRequestError(err.Error())
	}

	fileReader, err := file.Open()
	if err != nil {
		return apperror.BadRequestError(err.Error())
	}
	defer fileReader.Close()

	dto := UploadFileDTO{
		Id:     time.Now().UnixMilli(),
		Size:   file.Size,
		Reader: fileReader,
	}

	if err := h.ImageService.UploadImage(dto); err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "ok!",
		"id":      dto.Id,
	})
}

package image

import (
	"encoding/json"
	"net/http"
	"producer/internal/apperror"

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
	router.HandlerFunc(http.MethodPost, imageDownloadURL, apperror.Middleware(h.DownloadImage))
}

func (h *Handler) UploadImage(w http.ResponseWriter, r *http.Request) error {

	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		return apperror.BadRequestError(err.Error())
	}
	defer file.Close()

	if err := h.ImageService.UploadImage(UploadDTO{FileHeader: fileHeader, File: file}); err != nil {
		return apperror.BadRequestError(err.Error())
	}

	return json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "ok!",
		"id":      fileHeader.Filename,
	})
}

func (h *Handler) DownloadImage(w http.ResponseWriter, r *http.Request) error {
	return h.ImageService.DownloadImage(0)
}

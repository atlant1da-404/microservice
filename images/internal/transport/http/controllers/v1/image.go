package v1

import (
	"github.com/julienschmidt/httprouter"
	"images/internal/service"
	"net/http"
)

const (
	imageUploadURL   = "/api/image/upload"
	imageDownloadURL = "/api/image/download/:id"
)

type ImageHandler struct {
	ImageService service.ImageService
}

func (h *ImageHandler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, imageUploadURL, h.UploadImage)
	router.HandlerFunc(http.MethodGet, imageDownloadURL, h.DownloadImage)
}

func (h *ImageHandler) UploadImage(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *ImageHandler) DownloadImage(w http.ResponseWriter, r *http.Request) {
	return
}

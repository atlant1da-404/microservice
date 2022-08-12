package image

import (
	"encoding/json"
	"net/http"
	"producer/internal/apperror"
	"strconv"

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
	router.HandlerFunc(http.MethodGet, imageDownloadURL, apperror.Middleware(h.DownloadImage))
}

func (h *Handler) UploadImage(w http.ResponseWriter, r *http.Request) error {

	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		return apperror.BadRequestError(err.Error())
	}
	defer file.Close()

	id, err := h.ImageService.UploadImage(UploadDTO{FileHeader: fileHeader, File: file})
	if err != nil {
		return apperror.BadRequestError(err.Error())
	}

	return json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "ok!",
		"id":      id,
	})
}

func (h *Handler) DownloadImage(w http.ResponseWriter, r *http.Request) error {

	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)

	imageId, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		return apperror.BadRequestError(err.Error())
	}

	quality := r.URL.Query().Get("quality")
	img, err := h.ImageService.DownloadImage(DownloadDTO{Id: int64(imageId), Quality: quality})
	if err != nil {
		return apperror.BadRequestError(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	_, err = w.Write([]byte(img.Base64))
	if err != nil {
		return apperror.BadRequestError(err.Error())
	}

	return json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "ok!",
		"id":      imageId,
		"quality": quality,
	})
}

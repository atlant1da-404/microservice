package image

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"math/rand"
	"net/http"
	"producer/internal/apperror"
	"producer/pkg/rest"
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
		Id:     rand.Int(),
		Size:   file.Size,
		Reader: fileReader,
	}

	if err := h.ImageService.UploadImage(r.Context(), dto); err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "ok!",
		"id":      dto.Id,
	})
}

func (h *Handler) DownloadImage(w http.ResponseWriter, r *http.Request) error {

	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)

	dto := DownloadFileDTO{
		ID:      params.ByName("id"),
		Quality: r.URL.Query().Get("quality"),
	}

	img, err := h.ImageService.DownloadImage(r.Context(), dto)
	if err != nil {
		return err
	}

	return rest.SendFile(w, r, img.ID, img.Bytes)
}

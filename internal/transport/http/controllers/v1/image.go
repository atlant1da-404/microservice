package v1

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"images/internal/model"
	"images/internal/service"
	"images/pkg/apperror"
	"images/pkg/rest"
	"math/rand"
	"net/http"
)

const (
	imageUploadURL   = "/api/v1/image/upload"
	imageDownloadURL = "/api/v1/image/download/:id"
)

type ImageHandler struct {
	ImageService service.ImageService
}

func NewImageHandler(imageService service.ImageService) *ImageHandler {
	return &ImageHandler{ImageService: imageService}
}

func (h *ImageHandler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, imageUploadURL, apperror.Middleware(h.UploadImage))
	router.HandlerFunc(http.MethodGet, imageDownloadURL, apperror.Middleware(h.DownloadImage))
}

// UploadImage api endpoint to save image from client.
// Use HTTP protocol.
// REST.
// Endpoint: /api/v1/image/upload .
// Method: POST.
// Queries: none, params: none.
// Returns: id and confirmation message or errors.
func (h *ImageHandler) UploadImage(w http.ResponseWriter, r *http.Request) error {

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

	dto := model.UploadFileDTO{
		Id:     rand.Int(),
		Size:   file.Size,
		Reader: fileReader,
	}

	if err := h.ImageService.CreateImage(r.Context(), dto); err != nil {
		return err
	}

	return json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "ok!",
		"id":      dto.Id,
	})
}

// DownloadImage api endpoint to download image from server.
// Use HTTP protocol.
// REST.
// Endpoint: /api/v1/image/download/:id .
// Method: GET.
// Queries: yes, params: yes.
// Returns: file (image) in bytes to client.
func (h *ImageHandler) DownloadImage(w http.ResponseWriter, r *http.Request) error {

	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)

	dto := model.DownloadFileDTO{
		Id:      params.ByName("id"),
		Quality: r.URL.Query().Get("quality"),
	}

	img, err := h.ImageService.GetImage(r.Context(), dto)
	if err != nil {
		return err
	}

	return rest.SendFile(w, r, img.ID, img.Bytes)
}

package resize

import (
	"image"
)

// Compressor is interface package for compress images
type Compressor interface {
	// GetImage returns image for bytes
	GetImage(data []byte) (image.Image, error)
	// Compress returns a slice of images with right condition of calculation
	Compress(img image.Image) ([]image.Image, []int)
	// GeneratePictureId returns id for image from his unical id and quality
	GeneratePictureId(formId string, quality int) string
}

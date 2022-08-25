package resize

import (
	"bytes"
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
)

const stepOptimization = 3

// Compressor is interface package for compress images
type Compressor interface {
	// GetImage returns image for bytes
	GetImage(data []byte) (image.Image, error)
	// Compress returns a slice of images with right condition of calculation
	Compress(img image.Image) ([]image.Image, []int)
	// GeneratePictureId returns id for image from his unical id and quality
	GeneratePictureId(formId string, quality int) string
	// reSize generate a images from width and height
	reSize(img image.Image, width, height []uint) (pictures []image.Image)
}

func (c *Resizer) GetImage(data []byte) (image.Image, error) {

	img, err := jpeg.Decode(bytes.NewReader(data))
	if err == nil {
		return img, nil
	}

	img, err = png.Decode(bytes.NewReader(data))
	if err == nil {
		return img, nil
	}

	return nil, err
}

func (c *Resizer) Compress(img image.Image) ([]image.Image, []int) {

	photoHeight := uint(img.Bounds().Size().Y)
	photoWidth := uint(img.Bounds().Size().X)

	width := []uint{photoWidth, photoWidth - (photoWidth / 4), photoWidth / 2, photoWidth / 4}
	height := []uint{photoHeight, photoHeight - (photoHeight / 4), photoHeight / 2, photoHeight / 4}
	quality := []int{100, 75, 50, 25}

	return c.reSize(img, width, height), quality
}

func (c *Resizer) GeneratePictureId(formId string, quality int) string {
	return fmt.Sprintf("%s_%d.jpeg", formId, quality)
}

func (c *Resizer) reSize(img image.Image, width, height []uint) (pictures []image.Image) {

	for i := 0; i <= stepOptimization; i++ {
		pictures = append(pictures, resize.Resize(width[i], height[i], img, resize.Lanczos3))
	}

	return pictures
}

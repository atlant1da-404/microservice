package convert

import (
	"encoding/base64"
	"image"
	"strings"
)

func Base64Dec(base64Enc string) (image.Image, error) {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64Enc))
	img, _, err := image.Decode(reader)
	return img, err
}

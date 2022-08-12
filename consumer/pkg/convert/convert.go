package convert

import (
	"bytes"
	"encoding/base64"
	"image"
	"io"
	"mime/multipart"
	"strings"
)

func Base64Dec(base64Enc string) (image.Image, error) {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64Enc))
	img, _, err := image.Decode(reader)
	return img, err
}

func Base64Enc(file multipart.File) (string, error) {
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

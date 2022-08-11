package convert

import (
	"bytes"
	"encoding/base64"
	"io"
	"mime/multipart"
)

func Base64Enc(file multipart.File) (string, error) {

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

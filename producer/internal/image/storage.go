package image

type Storage interface {
	UploadImage(bData []byte) error
}

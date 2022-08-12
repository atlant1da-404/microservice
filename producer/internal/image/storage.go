package image

type StorageAmqp interface {
	UploadImage(bFile []byte) error
}

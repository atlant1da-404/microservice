package image

type StorageAmqp interface {
	UploadImage(bData []byte) error
	DownloadImage(bData []byte) ([]byte, error)
}

type StorageCache interface {
	CheckInCache(id int64) bool
}

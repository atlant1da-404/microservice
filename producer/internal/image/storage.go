package image

type StorageAmqp interface {
	UploadImage(bFile []byte) error
}

type StorageMinio interface {
	DownloadImage(fileId string) (*File, error)
}

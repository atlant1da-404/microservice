package storage

type ImageStorage interface {
	DownloadFile()
	UploadFile()
}

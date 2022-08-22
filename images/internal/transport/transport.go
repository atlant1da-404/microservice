package transport

import "net/http"

type ImageHandler interface {
	UploadImage(w http.ResponseWriter, r *http.Request)
	DownloadImage(w http.ResponseWriter, r *http.Request)
}

type AmqpHandler interface {
	Producer()
	Consumer()
}

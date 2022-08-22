package v1

import "images/internal/service"

type AmqpHandler struct {
	ImageService service.ImageService
}

func (a *AmqpHandler) Producer() {
	return
}

func (a *AmqpHandler) Consumer() {
	return
}

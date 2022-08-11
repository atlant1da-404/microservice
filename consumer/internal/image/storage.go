package image

type Storage interface {
	Set(modelId int) error
}

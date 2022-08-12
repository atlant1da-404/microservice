package image

type File struct {
	ID    string `json:"id"`
	Size  int64  `json:"size"`
	Bytes []byte `json:"file"`
}

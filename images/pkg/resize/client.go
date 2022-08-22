package resize

type compressor struct {
}

func NewCompressor() Compressor {
	return &compressor{}
}

package interf

type CompressOp interface {
	Operation
	Compress(data []byte) ([]byte, error)
	Decompress(data []byte) ([]byte, error)
}

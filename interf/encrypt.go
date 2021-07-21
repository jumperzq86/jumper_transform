package interf

type EncryptOp interface {
	Operation
	Encrypt(data []byte) ([]byte, error)
	Decrypt(data []byte) ([]byte, error)
}

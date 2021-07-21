package interf

type Message struct {
	Type    uint16
	Content []byte
}

type PacketOp interface {
	Operation
	Pack(originData interface{}) ([]byte, error)
	Unpack(packData []byte, obj interface{}) error
}

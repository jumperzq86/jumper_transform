package packet_test

import (
	"github.com/jumperzq86/jumper_transform/def"
	"github.com/jumperzq86/jumper_transform/interf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/golang/protobuf/proto"
	. "github.com/jumperzq86/jumper_transform/impl/packet"
)

var _ = Describe("PacketProtobuf", func() {

	var op interf.Operation
	BeforeEach(func() {
		op = NewpacketOpProtobuf(nil)
	})

	Describe("test protobuf", func() {
		Context("", func() {
			bodyData := "test protobuf"
			info := &StringMessage{
				Body: proto.String(bodyData),
				Header: &Header{
					MessageId: proto.String("20-05"),
					Topic:     proto.String("golang"),
				},
			}

			pack := []byte{0x0a, 0x0f, 0x0a, 0x05, 0x32, 0x30, 0x2d, 0x30, 0x35, 0x12, 0x06, 0x67, 0x6f, 0x6c, 0x61,
				0x6e, 0x67, 0x12, 0x0d, 0x74, 0x65, 0x73, 0x74, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66}

			It("", func() {
				var output []byte
				rst, err := op.Operate(def.Forward, info, &output)
				Expect(rst).To(Equal(true))
				BeNil().Match(err)
				Expect(output).To(Equal(pack))
			})

			It("", func() {
				var output StringMessage
				rst, err := op.Operate(def.Backward, pack, &output)
				Expect(rst).To(Equal(true))
				BeNil().Match(err)
				Expect(output).To(Equal(*info))

			})
		})

	})
})

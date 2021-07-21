package packet_test

import (
	"github.com/jumperzq86/jumper_transform/def"
	"github.com/jumperzq86/jumper_transform/interf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	//
	. "github.com/jumperzq86/jumper_transform/impl/packet"
)

var _ = Describe("PacketBinary", func() {
	var op interf.Operation

	BeforeEach(func() {
		op = NewpacketOpBinary(nil)
	})

	Describe("test binary", func() {
		Context("packet with binary", func() {
			info := &interf.Message{
				Type:    99,
				Content: []byte("test binary"),
			}
			pack := []byte{0x00, 0x63, 0x74, 0x65, 0x73,
				0x74, 0x20, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79}

			It("", func() {
				var output []byte
				rst, err := op.Operate(def.Forward, info, &output)
				Expect(rst).To(Equal(true))
				BeNil().Match(err)
				Expect(output).To(Equal(pack))
			})

			It("", func() {
				var output interf.Message
				rst, err := op.Operate(def.Backward, pack, &output)
				Expect(rst).To(Equal(true))
				BeNil().Match(err)
				Expect(output).To(Equal(*info))
			})
		})

	})
})

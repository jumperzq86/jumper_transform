package packet_test

import (
	"github.com/jumperzq86/jumper_transform/def"
	"github.com/jumperzq86/jumper_transform/interf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	//
	. "github.com/jumperzq86/jumper_transform/impl/packet"
)

var _ = Describe("PacketBase64", func() {

	var op interf.Operation
	BeforeEach(func() {
		op = NewpacketOpBase64(nil)
	})

	Describe("test base64", func() {
		Context("", func() {
			info := []byte("test base64")
			pack := []byte("dGVzdCBiYXNlNjQ=")

			It("", func() {
				var output []byte
				rst, err := op.Operate(def.Forward, info, &output)
				Expect(rst).To(Equal(true))
				BeNil().Match(err)
				Expect(output).To(Equal(pack))
			})

			It("", func() {
				var output []byte
				rst, err := op.Operate(def.Backward, pack, &output)
				Expect(rst).To(Equal(true))
				BeNil().Match(err)
				Expect(output).To(Equal(info))

			})
		})

		Context("", func() {
			info := []byte("input: test base64")
			pack := []byte("aW5wdXQ6IHRlc3QgYmFzZTY0")

			It("packet", func() {
				var output []byte
				rst, err := op.Operate(def.Forward, info, &output)
				Expect(rst).To(Equal(true))
				BeNil().Match(err)
				Expect(output).To(Equal(pack))
			})

			It("unpacket", func() {
				var output []byte
				rst, err := op.Operate(def.Backward, pack, &output)
				Expect(rst).To(Equal(true))
				BeNil().Match(err)
				Expect(output).To(Equal(info))

			})
		})
	})
})

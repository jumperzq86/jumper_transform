package packet_test

import (
	"github.com/jumperzq86/jumper_transform/def"
	"github.com/jumperzq86/jumper_transform/interf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	//
	. "github.com/jumperzq86/jumper_transform/impl/packet"
)

var _ = Describe("PacketBase64Url", func() {

	var op interf.Operation
	BeforeEach(func() {
		op = NewpacketOpBase64RawUrl(nil)
	})

	Describe("test base64 raw url", func() {
		Context("", func() {
			info := []byte("https://duckduckgo.com/?q=golang+gorilla")
			pack := []byte("aHR0cHM6Ly9kdWNrZHVja2dvLmNvbS8_cT1nb2xhbmcrZ29yaWxsYQ")

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

	})
})

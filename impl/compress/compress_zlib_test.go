package compress_test

import (
	"github.com/jumperzq86/jumper_transform/def"
	"github.com/jumperzq86/jumper_transform/interf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strings"

	. "github.com/jumperzq86/jumper_transform/impl/compress"
)

var _ = Describe("CompressZlib", func() {

	var op interf.Operation
	BeforeEach(func() {
		op = NewcompressOpZlib(nil)
	})

	Describe("test zlib", func() {
		Context("", func() {
			info := []byte(strings.Repeat("test zlib", 100))
			comp := []byte{0x78, 0x9c, 0x2a, 0x49, 0x2d, 0x2e, 0x51, 0xa8, 0xca, 0xc9, 0x4c, 0x1a, 0x65, 0x8c, 0x32,
				0x46, 0x19, 0x03, 0xc2, 0x00, 0x04, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x6b, 0x64, 0xb4}
			It("", func() {
				var output []byte
				rst, err := op.Operate(def.Forward, info, &output)
				Expect(rst).To(Equal(true))
				BeNil().Match(err)
				Expect(output).To(Equal(comp))
			})

			It("", func() {
				var output []byte
				rst, err := op.Operate(def.Backward, comp, &output)
				Expect(rst).To(Equal(true))
				BeNil().Match(err)
				Expect(output).To(Equal(info))
			})
		})
	})
})

package compress_test

import (
	"github.com/jumperzq86/jumper_transform/def"
	. "github.com/jumperzq86/jumper_transform/impl/compress"
	"github.com/jumperzq86/jumper_transform/interf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strings"
	//"strings"
)

var _ = Describe("CompressGzip", func() {

	var op interf.Operation
	BeforeEach(func() {
		op = NewcompressOpGzip(nil)
	})

	Describe("test gzip", func() {
		Context("", func() {
			info := []byte(strings.Repeat("test gzip", 100))
			comp := []byte{0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x2a, 0x49, 0x2d, 0x2e, 0x51, 0x48,
				0xaf, 0xca, 0x2c, 0x18, 0x65, 0x8c, 0x32, 0x46, 0x19, 0x03, 0xc2, 0x00, 0x04, 0x00, 0x00, 0xff, 0xff,
				0x0c, 0xde, 0x3d, 0xcf, 0x84, 0x03, 0x00, 0x00}
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

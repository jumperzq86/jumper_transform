package encrypt_test

import (
	"github.com/jumperzq86/jumper_transform/def"
	"github.com/jumperzq86/jumper_transform/interf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/jumperzq86/jumper_transform/impl/encrypt"
)

var _ = Describe("EncryptMd5", func() {

	var op interf.Operation
	BeforeEach(func() {
		op = NewencryptOpMd5(nil)
	})

	Describe("test md5", func() {
		Context("", func() {
			info := []byte("test md5")
			ency := []byte{0x0e, 0x4e, 0x3b, 0x26, 0x81, 0xe8, 0x93, 0x1c, 0x06, 0x7a, 0x23, 0xc5, 0x83, 0xc8, 0x78, 0xd5}
			It("", func() {
				var output []byte
				rst, err := op.Operate(def.Forward, info, &output)
				Expect(rst).To(Equal(true))
				BeNil().Match(err)
				Expect(output).To(Equal(ency))
			})
		})
	})
})

package encrypt_test

import (
	"github.com/jumperzq86/jumper_transform/def"
	"github.com/jumperzq86/jumper_transform/interf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/jumperzq86/jumper_transform/impl/encrypt"
)

var _ = Describe("EncryptSha256", func() {

	var op interf.Operation
	BeforeEach(func() {
		op = NewencryptOpSha256(nil)
	})

	Describe("test sha256", func() {
		Context("", func() {
			info := []byte("test sha256")
			ency := []byte{0xc7, 0x1d, 0x13, 0x7d, 0xa1, 0x40, 0xc5, 0xaf, 0xef, 0xd7, 0xdb, 0x8e, 0x7a, 0x25, 0x5d,
				0xf4, 0x5c, 0x2a, 0xc4, 0x60, 0x64, 0xe9, 0x34, 0x41, 0x6d, 0xc0, 0x40, 0x20, 0xa9, 0x1f, 0x3f, 0xd2}
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

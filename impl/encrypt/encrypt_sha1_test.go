package encrypt_test

import (
	"github.com/jumperzq86/jumper_transform/def"
	. "github.com/jumperzq86/jumper_transform/impl/encrypt"
	"github.com/jumperzq86/jumper_transform/interf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EncryptSha1", func() {

	var op interf.Operation
	BeforeEach(func() {
		op = NewencryptOpSha1(nil)
	})

	Describe("test sha1", func() {
		Context("", func() {
			info := []byte("test sha1")
			ency := []byte{0xb9, 0x9c, 0x07, 0x13, 0x33, 0xd4, 0xdb, 0xca, 0x0d,
				0x92, 0x98, 0xe5, 0xc8, 0xd7, 0x48, 0x0f, 0x17, 0x6c, 0xaf, 0xdc}
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

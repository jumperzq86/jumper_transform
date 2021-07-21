package encrypt_test

import (
	"github.com/jumperzq86/jumper_transform/def"
	. "github.com/jumperzq86/jumper_transform/impl/encrypt"
	"github.com/jumperzq86/jumper_transform/interf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EncryptAes", func() {

	var op interf.Operation
	BeforeEach(func() {
		op = NewencryptOpAes([]interface{}{[]byte("abcdefghijklmnop")})
	})

	//note: 此处因为加密过程中使用的 iv 需要随机产生，因此即便同样的数据和密钥，每次加密之后的结果都是不同的。因此无法对中间结果进行校验
	Describe("test aes", func() {
		Context("", func() {
			info := []byte("test aes")

			It("", func() {
				var ency []byte
				var output []byte
				rst, err := op.Operate(def.Forward, info, &ency)
				Expect(rst).To(Equal(true))
				BeNil().Match(err)

				rst, err = op.Operate(def.Backward, ency, &output)
				Expect(rst).To(Equal(true))
				BeNil().Match(err)

				Expect(output).To(Equal(info))
			})
		})
	})
})

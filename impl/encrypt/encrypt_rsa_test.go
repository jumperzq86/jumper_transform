package encrypt_test

import (
	"github.com/jumperzq86/jumper_transform/interf"

	"github.com/jumperzq86/jumper_transform/def"
	. "github.com/jumperzq86/jumper_transform/impl/encrypt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EncryptRsa", func() {
	var op interf.Operation
	var privKeyLocal []byte
	var pubKeyRemote []byte
	BeforeEach(func() {
		privKeyLocal = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgHj092QwqywUYwcs/IalZgbLb5/CDIBkAMVQENIO1bxE53ddf7Ex
qQ82OxAH+2K+5wVZSJAM4Gsu8k+bBntRAdZtxODAj8XnMqlrU06bzBcZA1HBta79
wBAWcT2z8HL6n4BUXm1ZVilsOemN+YhZmaoqZjSTSIUUP0dD8fxl2FSfAgMBAAEC
gYAnz+F5vbcpjBBINVts6hXZ2K4F9HXu8Ht8dm0C5tUc4cUZ+dFkvka59gycqzPn
/ZPGo+uJVmOrg8lHwGvyvOgQeZYMpozHx9erVBjkFGa0zl7N8cOr26tSd1WjOTsz
pbdpNdbwUvuQRxtLUtSoUi9qBouX9/SMsBInxoKnElbKeQJBANBDW5J3vBOB8Eu1
yoB2EGZKfNCnj1jad1dvKRUtRdZsajhoPODLjOvTk+p0yzSPNE5cKs64n6qv5DSe
Qz3cGdUCQQCUrpRT/pcFvkk2io+jiIVo2YMPG+UMgCAgQ24ImyvlcSffUMYyHo1T
mptlZxpqXR30kAT0cUUGEwa4eTfzo1qjAkEAmohwiVUJHFx9ZmBFG17/ujcpOFYs
2px4k0sra8hFGt8lTieijJbdQRuklQMGyUA9aqufI/cAmtGa/qQ0dtE75QJAGuuJ
G7Qq7/VGGEP6QefED/niCuhIDn1cU7shvxkpV3+ncn2ThRuXBx4lKLUESWHWvKMe
4otnrfPf8aKxCMH49QJANyi1LmbuuMr8jWgpEV0ZutEYiCclK8hEAuEsqbDuvagg
GuAeT9vSgKpBXkUodUbuQABhp6O1aDNWFyksoYa7jg==
-----END RSA PRIVATE KEY-----
`)

		pubKeyRemote = []byte(`
-----BEGIN PUBLIC KEY-----
MIGeMA0GCSqGSIb3DQEBAQUAA4GMADCBiAKBgHj092QwqywUYwcs/IalZgbLb5/C
DIBkAMVQENIO1bxE53ddf7ExqQ82OxAH+2K+5wVZSJAM4Gsu8k+bBntRAdZtxODA
j8XnMqlrU06bzBcZA1HBta79wBAWcT2z8HL6n4BUXm1ZVilsOemN+YhZmaoqZjST
SIUUP0dD8fxl2FSfAgMBAAE=
-----END PUBLIC KEY-----
`)
		op = NewencryptOpRsa([]interface{}{pubKeyRemote, privKeyLocal})
	})

	//note: 此处因为加密过程中使用随机数，因此即便同样的数据和密钥，每次加密之后的结果都是不同的。因此无法对中间结果进行校验
	Describe("test rsa", func() {
		Context("", func() {
			info := []byte("test rsa")
			It("", func() {
				var output []byte
				var ency []byte

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

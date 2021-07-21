package encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"github.com/jumperzq86/jumper_transform/interf"

	"github.com/jumperzq86/jumper_transform/def"

	"github.com/jumperzq86/jumper_transform/util"
)

type encryptOpRsa struct {
	rsaPublicKeyRemote []byte // 来自对端生成的公钥，用于加密
	rsaPrivateKeyLocal []byte // 来自本端生成的私钥
}

func NewencryptOpRsa(params []interface{}) interf.EncryptOp {
	var op encryptOpRsa
	op.init(params)
	return &op
}

func (self *encryptOpRsa) init(params []interface{}) bool {
	if len(params) != 2 {

		return false
	}

	var ok bool
	self.rsaPublicKeyRemote, ok = params[0].([]byte)
	if !ok {

		return false
	}
	self.rsaPrivateKeyLocal, ok = params[1].([]byte)
	if !ok {

		return false
	}

	if len(self.rsaPrivateKeyLocal) > 2048 {

		return false
	}

	return true
}

func (self *encryptOpRsa) Operate(direct int8, input interface{}, output interface{}) (bool, error) {

	if direct == def.Forward {
		tmpOutput, err := self.Encrypt(input.([]byte))
		if err != nil {

			return false, err
		}
		*(output.(*[]byte)) = tmpOutput
		return true, nil

	} else {
		tmpOutput, err := self.Decrypt(input.([]byte))
		if err != nil {

			return false, err
		}
		*(output.(*[]byte)) = tmpOutput
		return true, nil
	}

	return true, nil
}

func (self *encryptOpRsa) Encrypt(data []byte) ([]byte, error) {

	defer util.TraceLog("encryptOpRsa.Encrypt")()
	block, _ := pem.Decode(self.rsaPublicKeyRemote)
	if block == nil {
		return nil, def.ErrInvalidRsaPublicKey

	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, data)
}

func (self *encryptOpRsa) Decrypt(data []byte) ([]byte, error) {

	defer util.TraceLog("encryptOpRsa.Decrypt")()
	block, _ := pem.Decode(self.rsaPrivateKeyLocal)
	if block == nil {
		return nil, def.ErrInvalidRsaPrivateKey

	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err

	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, data)
}

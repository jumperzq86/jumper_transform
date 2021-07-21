package encrypt

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"io"

	"github.com/jumperzq86/jumper_transform/interf"

	"github.com/jumperzq86/jumper_transform/def"

	"github.com/jumperzq86/jumper_transform/util"
)

type encryptOpDes struct {
	desKey []byte
}

func NewencryptOpDes(params []interface{}) interf.EncryptOp {
	var op encryptOpDes
	op.init(params)
	return &op
}

func (self *encryptOpDes) init(params []interface{}) bool {

	if params == nil || len(params) != 1 {

		return false
	}

	var ok bool
	self.desKey, ok = params[0].([]byte)
	if !ok {

		return false
	}

	if len(self.desKey) != 64/8 {

		return false
	}

	return true
}

func (self *encryptOpDes) Operate(direct int8, input interface{}, output interface{}) (bool, error) {

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

func (self *encryptOpDes) Encrypt(data []byte) ([]byte, error) {

	defer util.TraceLog("encryptOpDes.Encrypt")()
	if self.desKey == nil {
		return nil, def.ErrInvalidDesKey
	}
	if data == nil {
		return nil, def.ErrParamShouldNotNil
	}

	block, err := des.NewCipher(self.desKey)
	if err != nil {
		panic(err)
	}

	//填充字节
	data = pKCS5Padding(data, block.BlockSize())

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, des.BlockSize+len(data))
	iv := ciphertext[:des.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[des.BlockSize:], data)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.

	//
	return ciphertext, nil
}

func (self *encryptOpDes) Decrypt(data []byte) ([]byte, error) {

	defer util.TraceLog("encryptOpDes.Decrypt")()
	block, err := des.NewCipher(self.desKey)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(data) < des.BlockSize {
		panic("ciphertext too short")

	}
	iv := data[:des.BlockSize]
	data = data[des.BlockSize:]

	// CBC mode always works in whole blocks.
	if len(data)%des.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")

	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(data, data)

	data = pKCS5UnPadding(data)

	//
	return data, nil
}

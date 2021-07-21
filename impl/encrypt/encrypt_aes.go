package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"github.com/jumperzq86/jumper_transform/interf"

	"github.com/jumperzq86/jumper_transform/def"

	"github.com/jumperzq86/jumper_transform/util"
)

type encryptOpAes struct {
	aesKey []byte
}

func NewencryptOpAes(params []interface{}) interf.EncryptOp {
	var op encryptOpAes
	op.init(params)
	return &op
}

func (self *encryptOpAes) init(params []interface{}) bool {

	if params == nil || len(params) != 1 {

		return false
	}

	var ok bool
	self.aesKey, ok = params[0].([]byte)
	if !ok {

		return false
	}

	if len(self.aesKey) != 128/8 && len(self.aesKey) != 192/8 && len(self.aesKey) != 256/8 {

		return false
	}

	return true
}

func (self *encryptOpAes) Operate(direct int8, input interface{}, output interface{}) (bool, error) {

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

func (self *encryptOpAes) Encrypt(data []byte) ([]byte, error) {

	defer util.TraceLog("encryptOpAes.Encrypt")()
	if self.aesKey == nil {
		return nil, def.ErrInvalidAesKey
	}
	if data == nil {
		return nil, def.ErrParamShouldNotNil
	}

	block, err := aes.NewCipher(self.aesKey)
	if err != nil {
		panic(err)
	}

	//填充字节
	data = pKCS5Padding(data, block.BlockSize())

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], data)

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.

	//
	return ciphertext, nil
}

func (self *encryptOpAes) Decrypt(data []byte) ([]byte, error) {

	defer util.TraceLog("encryptOpAes.Decrypt")()
	block, err := aes.NewCipher(self.aesKey)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(data) < aes.BlockSize {
		panic("ciphertext too short")

	}
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]

	// CBC mode always works in whole blocks.
	if len(data)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")

	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(data, data)

	data = pKCS5UnPadding(data)

	// If the original plaintext lengths are not a multiple of the block
	// size, padding would have to be added when encrypting, which would be
	// removed at self point. For an example, see
	// https://tools.ietf.org/html/rfc5246#section-6.2.3.2. However, it's
	// critical to note that ciphertexts must be authenticated (i.e. by
	// using crypto/hmac) before being decrypted in order to avoid creating
	// a padding oracle.

	//
	return data, nil
}

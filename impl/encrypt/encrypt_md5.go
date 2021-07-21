package encrypt

import (
	"crypto/md5"

	"github.com/jumperzq86/jumper_transform/interf"

	"github.com/jumperzq86/jumper_transform/def"

	"github.com/jumperzq86/jumper_transform/util"
)

type encryptOpMd5 struct {
}

func NewencryptOpMd5(params []interface{}) interf.EncryptOp {
	var op encryptOpMd5
	op.init(params)
	return &op
}

func (self *encryptOpMd5) init(params []interface{}) bool {
	return true
}

func (self *encryptOpMd5) Operate(direct int8, input interface{}, output interface{}) (bool, error) {

	if direct == def.Forward {
		tmpOutput, err := self.Encrypt(input.([]byte))
		if err != nil {

			return false, err
		}
		*(output.(*[]byte)) = tmpOutput
		return true, nil

	} else {

		return false, def.ErrMd5NoDecrypt
	}

	return true, nil
}

func (*encryptOpMd5) Encrypt(data []byte) ([]byte, error) {
	defer util.TraceLog("encryptOpMd5.Encrypt")()
	r := md5.Sum(data)
	rst := r[:]
	return rst, nil

}

func (*encryptOpMd5) Decrypt(data []byte) ([]byte, error) {
	defer util.TraceLog("encryptOpMd5.Decrypt")()
	return nil, def.ErrMd5NoDecrypt
}

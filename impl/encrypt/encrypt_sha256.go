package encrypt

import (
	"crypto/sha256"

	"github.com/jumperzq86/jumper_transform/interf"

	"github.com/jumperzq86/jumper_transform/def"

	"github.com/jumperzq86/jumper_transform/util"
)

type encryptOpSha256 struct {
}

func NewencryptOpSha256(params []interface{}) interf.EncryptOp {
	var op encryptOpSha256
	op.init(params)
	return &op
}

func (self *encryptOpSha256) init(params []interface{}) bool {
	return true
}

func (self *encryptOpSha256) Operate(direct int8, input interface{}, output interface{}) (bool, error) {

	if direct == def.Forward {
		tmpOutput, err := self.Encrypt(input.([]byte))
		if err != nil {

			return false, err
		}
		*(output.(*[]byte)) = tmpOutput
		return true, nil

	} else {

		return false, def.ErrSha256NoDecrypt
	}

	return true, nil
}

func (*encryptOpSha256) Encrypt(data []byte) ([]byte, error) {
	defer util.TraceLog("encryptOpSha256.Encrypt")()
	s := sha256.New()
	s.Write(data)
	r := s.Sum(nil)
	rst := r[:]
	return rst, nil

}

func (*encryptOpSha256) Decrypt(data []byte) ([]byte, error) {
	defer util.TraceLog("encryptOpSha256.Decrypt")()
	return nil, def.ErrSha1NoDecrypt
}

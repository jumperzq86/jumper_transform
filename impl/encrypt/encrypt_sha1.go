package encrypt

import (
	"crypto/sha1"

	"github.com/jumperzq86/jumper_transform/interf"

	"github.com/jumperzq86/jumper_transform/def"

	"github.com/jumperzq86/jumper_transform/util"
)

type encryptOpSha1 struct {
}

func NewencryptOpSha1(params []interface{}) interf.EncryptOp {
	var op encryptOpSha1
	op.init(params)
	return &op
}

func (self *encryptOpSha1) init(params []interface{}) bool {
	return true
}

func (self *encryptOpSha1) Operate(direct int8, input interface{}, output interface{}) (bool, error) {

	if direct == def.Forward {
		tmpOutput, err := self.Encrypt(input.([]byte))
		if err != nil {

			return false, err
		}
		*(output.(*[]byte)) = tmpOutput
		return true, nil

	} else {

		return false, def.ErrSha1NoDecrypt
	}

	return true, nil
}

func (*encryptOpSha1) Encrypt(data []byte) ([]byte, error) {
	defer util.TraceLog("encryptOpSha1.Encrypt")()
	s := sha1.New()
	s.Write(data)
	r := s.Sum(nil)
	rst := r[:]
	return rst, nil

}

func (*encryptOpSha1) Decrypt(data []byte) ([]byte, error) {
	defer util.TraceLog("encryptOpSha1.Decrypt")()
	return nil, def.ErrSha1NoDecrypt
}

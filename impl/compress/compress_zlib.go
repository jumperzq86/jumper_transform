package compress

import (
	"bytes"
	"compress/zlib"
	"io"

	"github.com/jumperzq86/jumper_transform/def"

	"github.com/jumperzq86/jumper_transform/interf"

	"github.com/jumperzq86/jumper_transform/util"
)

type compressOpZlib struct {
}

func NewcompressOpZlib(params []interface{}) interf.CompressOp {
	var op compressOpZlib
	op.init(params)
	return &op
}

func (self *compressOpZlib) init(params []interface{}) bool {
	return true
}

func (self *compressOpZlib) Operate(direct int8, input interface{}, output interface{}) (bool, error) {

	if direct == def.Forward {
		tmpOutput, err := self.Compress(input.([]byte))
		if err != nil {
			return false, err
		}
		*(output.(*[]byte)) = tmpOutput
		return true, nil

	} else {
		tmpOutput, err := self.Decompress(input.([]byte))
		if err != nil {
			return false, err
		}
		*(output.(*[]byte)) = tmpOutput
		return true, nil
	}

	return true, nil
}

func (self *compressOpZlib) Compress(data []byte) ([]byte, error) {
	defer util.TraceLog("compressOpZlib.Compress")()
	var buf bytes.Buffer
	c := zlib.NewWriter(&buf)

	_, err := c.Write(data)
	if err != nil {
		return nil, err
	}

	//note：若是上面使用　defer c.Close() 会导致在下面解压时出现错误：　decompress err: unexpected EOF, 使用c.Close则不会
	//但是对于decompress 中的　reader 却可以defer close
	c.Close()
	return buf.Bytes(), nil
}

func (self *compressOpZlib) Decompress(data []byte) ([]byte, error) {

	defer util.TraceLog("compressOpZlib.Decompress")()
	nr := bytes.NewReader(data)
	dc, err := zlib.NewReader(nr)
	if err != nil {
		return nil, err
	}
	defer dc.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, dc)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

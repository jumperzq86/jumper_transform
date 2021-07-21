package transform

import (
	"github.com/jumperzq86/jumper_transform/def"
	"github.com/jumperzq86/jumper_transform/interf"

	//"../compress"
	//"../encrypt"
	//"../packet"
	"github.com/jumperzq86/jumper_transform/impl/compress"
	"github.com/jumperzq86/jumper_transform/impl/encrypt"
	"github.com/jumperzq86/jumper_transform/impl/packet"
)

// 使用一系列　Operation 来构造一条操作链，
//　对于封包过程来说，　输入为　interface{}, 输出为　[]byte
//　对于解包过程来说，　输入为 []byte, 输出为 interface{}

//!!!note:　需要明确一点,　一个单独的PackageOpLink是支持并发调用的，即可以在多个协程中互不干扰的使用
// 这是因为　所有涉及到的数据只有两种情况
// 1. 入参传入，因此不同协程不会具有共用数据　2. 初始化的成员如direct或者密钥，他们都是初始化之后不再变化，即只进行读取，不进行写入，因此可以并发访问

//因此在实际项目使用中可以考虑定义一个全局的 map[string]Transform　不同的需求定义不同的　Transform　放入其中．

type Transform struct {
	opLink []interf.Operation
}

func (self *Transform) checkParam(opType int8, params []interface{}) bool {
	if opType <= def.PackageOpMin || opType >= def.PackageOpMax {

		return false
	}

	if (opType == def.EncryptAes || opType == def.EncryptDes) && len(params) != 1 {

		return false
	}

	if (opType == def.EncryptRsa) && len(params) != 2 {

		return false
	}

	return true
}

func (self *Transform) AddOp(opType int8, params []interface{}) bool {
	if !self.checkParam(opType, params) {
		return false
	}

	var op interf.Operation

	switch opType {
	//编码
	case def.PacketBase64:
		op = packet.NewpacketOpBase64(nil)
		break

	case def.PacketBase64RawUrl:
		op = packet.NewpacketOpBase64RawUrl(nil)
		break

	case def.PacketJson:
		op = packet.NewpacketOpJson(nil)
		break

	case def.PacketXml:
		op = packet.NewpacketOpXml(nil)
		break

	case def.PacketProtobuf:
		op = packet.NewpacketOpProtobuf(nil)
		break

	case def.PacketBinary:
		op = packet.NewpacketOpBinary(nil)
		break

		//压缩
	case def.CompressGzip:
		op = compress.NewcompressOpGzip(nil)
		break

	case def.CompressZlib:
		op = compress.NewcompressOpZlib(nil)
		break

		//加密
	case def.EncryptMd5:
		op = encrypt.NewencryptOpMd5(nil)
		break

	case def.EncryptSha1:
		op = encrypt.NewencryptOpSha1(nil)
		break

	case def.EncryptSha256:
		op = encrypt.NewencryptOpSha256(nil)
		break

	case def.EncryptAes:
		op = encrypt.NewencryptOpAes(params)
		break

	case def.EncryptDes:
		op = encrypt.NewencryptOpDes(params)
		break

	case def.EncryptRsa:
		op = encrypt.NewencryptOpRsa(params)
		break

	default:
		return false
		break
	}

	self.opLink = append(self.opLink, op)
	return true

}

func (self *Transform) Execute(direct int8, input interface{}, output interface{}) error {

	//这里是一个链式反应，因此需要根据op类型来构建中间类型
	//中间过程的输出都是 []byte
	var tmpOutput []byte
	var tmpInput interface{}
	tmpInput = input

	if direct == def.Forward {

		for index := 0; index <= len(self.opLink)-1; index++ {
			if index == len(self.opLink)-1 {
				rst, err := self.opLink[index].Operate(direct, tmpInput, output)
				if !rst {

					return err
				}

			} else {
				self.opLink[index].Operate(direct, tmpInput, &tmpOutput)

				tmpInput = tmpOutput
			}
		}

	} else {
		for index := len(self.opLink) - 1; index >= 0; index-- {
			if index == 0 {
				rst, err := self.opLink[index].Operate(direct, tmpInput, output)
				if !rst {

					return err
				}

			} else {
				self.opLink[index].Operate(direct, tmpInput, &tmpOutput)

				tmpInput = tmpOutput
			}
		}
	}

	return nil
}

func (self *Transform) Reset() {
	if len(self.opLink) != 0 {
		self.opLink = self.opLink[:0]
	}
	return
}

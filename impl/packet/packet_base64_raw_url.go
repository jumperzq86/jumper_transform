package packet

import (
	"encoding/base64"
	"reflect"

	"github.com/jumperzq86/jumper_transform/interf"

	"github.com/jumperzq86/jumper_transform/def"

	"github.com/jumperzq86/jumper_transform/util"
)

type packetOpBase64RawUrl struct {
}

func NewpacketOpBase64RawUrl(params []interface{}) interf.PacketOp {
	var op packetOpBase64RawUrl
	op.init(params)
	return &op
}

func (self *packetOpBase64RawUrl) init(params []interface{}) bool {
	return true
}

func (self *packetOpBase64RawUrl) Operate(direct int8, input interface{}, output interface{}) (bool, error) {

	if direct == def.Forward {
		tmpOutput, err := self.Pack(input)
		if err != nil {

			return false, err
		}
		*(output.(*[]byte)) = tmpOutput
		return true, nil

	} else {
		err := self.Unpack(input.([]byte), output)
		if err != nil {

			return false, err
		}
		return true, nil
	}

	return true, nil
}

//此函数中需要检查入参是否为 string / []byte
func (*packetOpBase64RawUrl) Pack(originData interface{}) ([]byte, error) {
	defer util.TraceLog("packetOpBase64RawUrl.Pack")()
	//需要判断入参类型为 string 或者 []byte
	vod := reflect.ValueOf(originData)
	tod := reflect.TypeOf(originData)
	if vod.IsValid() == false {

		return nil, def.ErrParamShouldNotNil
	}

	if vod.Kind() == reflect.String {
		rst := base64.RawURLEncoding.EncodeToString([]byte(originData.(string)))
		return []byte(rst), nil

	}

	if vod.Kind() == reflect.Slice && tod.Elem().Kind() == reflect.Uint8 {
		rst := base64.RawURLEncoding.EncodeToString(originData.([]byte))
		return []byte(rst), nil
	}

	return nil, def.ErrParamShouldStringOrBytes
}

func (*packetOpBase64RawUrl) Unpack(packData []byte, obj interface{}) error {

	defer util.TraceLog("packetOpBase64RawUrl.Unpack")()
	//判断接收结果的入参是一个*[]byte
	tod := reflect.TypeOf(obj)
	vod := reflect.ValueOf(obj)

	//这里需要注意区别 reflect.Value.Elem() 和 reflect.Type.Elem() 两个函数
	//要想查看 指针/数组/切片 等的元素类型应该使用 reflect.Type.Elem() 函数
	if vod.Kind() != reflect.Ptr || tod.Elem().Kind() != reflect.Slice || tod.Elem().Elem().Kind() != reflect.Uint8 {
		return def.ErrParamShouldPointOfByteSlice
	}

	rst, err := base64.RawURLEncoding.DecodeString(string(packData))
	if err != nil {
		return def.ErrDecodeFailed
	}

	//类型是指针，需要使用Elem获得指针指向的内存，然后进行值的设置
	vod.Elem().SetBytes(rst)
	return nil
}

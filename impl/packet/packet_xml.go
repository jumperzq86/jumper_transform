package packet

import (
	"encoding/xml"

	"github.com/jumperzq86/jumper_transform/def"

	"github.com/jumperzq86/jumper_transform/interf"

	"github.com/jumperzq86/jumper_transform/util"
)

type packetOpXml struct {
}

func NewpacketOpXml(params []interface{}) interf.PacketOp {
	var op packetOpXml
	op.init(params)
	return &op
}

func (self *packetOpXml) init(params []interface{}) bool {
	return true
}

func (self *packetOpXml) Operate(direct int8, input interface{}, output interface{}) (bool, error) {

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

//todo: xml是不能对 map 编码的, 这里需要添加检查
//https://stackoverflow.com/questions/30928770/marshall-map-to-xml-in-go?utm_medium=organic&utm_source=google_rich_qa&utm_campaign=google_rich_qa
func (*packetOpXml) Pack(originData interface{}) ([]byte, error) {
	defer util.TraceLog("packetOpXml.Pack")()
	return xml.Marshal(originData)
}

func (*packetOpXml) Unpack(packData []byte, obj interface{}) error {
	defer util.TraceLog("packetOpXml.Unpack")()

	//
	err := xml.Unmarshal(packData, obj)
	//
	return err
}

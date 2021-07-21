package packet

import (
	"encoding/json"
	"fmt"

	"github.com/jumperzq86/jumper_transform/def"

	"github.com/jumperzq86/jumper_transform/interf"

	"github.com/jumperzq86/jumper_transform/util"
)

type packetOpJson struct {
}

func NewpacketOpJson(params []interface{}) interf.PacketOp {
	var op packetOpJson
	op.init(params)
	return &op
}

func (self *packetOpJson) init(params []interface{}) bool {
	return true
}

func (self *packetOpJson) Operate(direct int8, input interface{}, output interface{}) (bool, error) {

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

func (*packetOpJson) Pack(originData interface{}) ([]byte, error) {
	defer util.TraceLog("packetOpJson.Pack")()

	return json.Marshal(originData)
}

func (*packetOpJson) Unpack(packData []byte, obj interface{}) error {
	defer util.TraceLog("packetOpJson.Unpack")()

	//关于解析动态内容：interface{} 参见如下网页：
	// http://cizixs.com/2016/12/19/golang-json-guide
	//
	err := json.Unmarshal(packData, obj)
	//

	return err

}

/////////////////////
//json.Unmarshal 函数的第二个参数可以是空接口的指针，而不是具体类型的指针
// 若是使用空接口的指针，那么在解析数据内容的时候，每一步都要进行类型断言，例子如下：
func test() {

	//切片
	strTest := "[{\"Title\":\"Casablanca\",\"released\": 1942,\"Actors\": [\"Humphrey Bogart\",\"Ingrid Bergman\"]},{\"Title\": \"Cool Hand Luke\",\"released\": 1967,\"color\": true,\"Actors\": [\"Paul Newman\"]}]"

	var tmptestun interface{}
	json.Unmarshal([]byte(strTest), &tmptestun)

	tmpstring := tmptestun.([]interface{})[0].(map[string]interface{})["Title"].(string)
	fmt.Println(tmpstring)
	//对象
	data := []byte(`{"Name":"cizixs","IsAdmin":true,"Followers":36}`)

	var f interface{}
	json.Unmarshal(data, &f)

	tmpName := f.(map[string]interface{})["Name"].(string)
	fmt.Println(tmpName)

}

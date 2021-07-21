package packet_test

import (
	"github.com/jumperzq86/jumper_transform/def"
	"github.com/jumperzq86/jumper_transform/interf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/jumperzq86/jumper_transform/impl/packet"
)

type XCard struct {
	Num  string `xml:"num"`
	Addr string `xml:"addr"`
}

type XmlStruct struct {
	Id    int64    `xml:"id"`
	Name  string   `xml:"name"`
	Age   int      `xml:"age"`
	Cards []*XCard `xml:"cards"`
}

var _ = Describe("PacketXml", func() {

	var op interf.Operation
	BeforeEach(func() {
		op = NewpacketOpXml(nil)
	})

	Describe("test xml", func() {
		Context("", func() {
			info := XmlStruct{
				Id:   1,
				Name: "wang",
				Age:  88,
				Cards: []*XCard{
					&XCard{
						Num:  "abcdefghijklmnopqrstuvwxyz",
						Addr: "beijing road",
					},
					&XCard{
						Num:  "012456789",
						Addr: "chengdu road",
					},
				},
			}
			pack := []byte("<XmlStruct><id>1</id><name>wang</name><age>88</age><cards><num>abcdefghijklmnopqrstuvwxyz</num>" +
				"<addr>beijing road</addr></cards><cards><num>012456789</num><addr>chengdu road</addr></cards></XmlStruct>")

			It("", func() {
				var output []byte
				rst, err := op.Operate(def.Forward, info, &output)
				Expect(rst).To(Equal(true))
				BeNil().Match(err)
				Expect(output).To(Equal(pack))
			})

			It("", func() {
				var output XmlStruct
				rst, err := op.Operate(def.Backward, pack, &output)
				Expect(rst).To(Equal(true))
				BeNil().Match(err)
				Expect(output).To(Equal(info))

			})
		})

	})
})

package packet_test

import (
	"github.com/jumperzq86/jumper_transform/def"
	"github.com/jumperzq86/jumper_transform/interf"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/jumperzq86/jumper_transform/impl/packet"
)

type Card struct {
	Num  string `json:"num"`
	Addr string `json:"addr"`
}

type JsonStruct struct {
	Id    int64   `json:"id"`
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	Cards []*Card `json:"cards"`
}

var _ = Describe("PacketJson", func() {

	var op interf.Operation
	BeforeEach(func() {
		op = NewpacketOpJson(nil)
	})

	Describe("test json", func() {
		Context("", func() {
			info := JsonStruct{
				Id:   1,
				Name: "wang",
				Age:  88,
				Cards: []*Card{
					&Card{
						Num:  "abcdefghijklmnopqrstuvwxyz",
						Addr: "beijing road",
					},
					&Card{
						Num:  "012456789",
						Addr: "chengdu road",
					},
				},
			}
			pack := []byte("{\"id\":1,\"name\":\"wang\",\"age\":88," +
				"\"cards\":[{\"num\":\"abcdefghijklmnopqrstuvwxyz\",\"addr\":\"beijing road\"}," +
				"{\"num\":\"012456789\",\"addr\":\"chengdu road\"}]}")

			It("", func() {
				var output []byte
				rst, err := op.Operate(def.Forward, info, &output)
				Expect(rst).To(Equal(true))
				BeNil().Match(err)
				Expect(output).To(Equal(pack))
			})

			It("", func() {
				var output JsonStruct
				rst, err := op.Operate(def.Backward, pack, &output)
				Expect(rst).To(Equal(true))
				BeNil().Match(err)
				Expect(output).To(Equal(info))

			})
		})

	})
})

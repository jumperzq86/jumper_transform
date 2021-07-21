package main

import (
	"fmt"

	"github.com/jumperzq86/jumper_transform"
	"github.com/jumperzq86/jumper_transform/def"
	"github.com/jumperzq86/jumper_transform/interf"

	"github.com/golang/protobuf/proto"
)

func main() {

	///////////////////////////////////////// 链接测试 /////////////////////////////////////////
	///////////////////////// json/xml ///////////////////////////
	PrivKeyLocal := []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgQCnCnuWcNacRnqwDfSNLx7bbJLJM+foyxqSzp/M0fYqjhMp8voe
51PUEGetCvM2kAakmRue6MXQ3TKrV7L6d3XTYGabBPzwDd0KoucklVVOS2vi1E7U
V1bZhB60YdayCb9dcnEdA0uyA+qQgk2VhMtP1fER8lll5EiUUT+T0vnq9wIDAQAB
AoGAfZo9Seb5CLNaR42GyK6Y1kdyrEYSaJJoHeGueTWbk24XbOCeQKSS/Q+E1bI5
JVrxE81o3nmLXT0mf35HP1yaRCrofCV7a4QBlD9CNkMfy68fJEA6gMFuVVAES6Fa
Zt1ENZ81NeENURUC+lLFSlUWm2Xbf+MZtCFIRE5Tj1HxvQkCQQDPTDZKpyqZ/1yg
PO1/Quu0iisDYROJMm4sHQowIYXkHA/pUQMEveomBGRLavWrN9t4oEotFAPi0qYW
847m7TmDAkEAzkkNyoz08+Dg4+SfwbjEyglyX7OkmOOGnCvEJldQm0wLZvrpJS6i
n24UiYx2Cg93BZrvD9Ce7oNEnwbnHG3yfQJAJtOce6ER3qQwwiaHSUXMhhU29zwQ
f6r9ba/Gv7sXq+EBre6phRLZL2O1MVcISph8t/w1yHmuPKa9yyC1TFV0ZwJADbeh
6SQybb04dy8OyI0G2QCD0IVbnqcSnnPymTIZNBp8b56jvks5mSxyxSrH9qdMnNzO
pNiUmPu1pnWJDMTq6QJAHIToUuuAN2z3pLpUJsM40T6sEwgbxiFPZ3iT4/T2Tgpy
BKLqQxR7jXKdl0iWYteC96pQ0bqytFse4lnmPMUCew==
-----END RSA PRIVATE KEY-----
`)

	PubKeyRemote := []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCnCnuWcNacRnqwDfSNLx7bbJLJ
M+foyxqSzp/M0fYqjhMp8voe51PUEGetCvM2kAakmRue6MXQ3TKrV7L6d3XTYGab
BPzwDd0KoucklVVOS2vi1E7UV1bZhB60YdayCb9dcnEdA0uyA+qQgk2VhMtP1fER
8lll5EiUUT+T0vnq9wIDAQAB
-----END PUBLIC KEY-----
`)

	type s1 struct {
		Name string
		Age  int
		Male bool
	}

	test1 := s1{
		Name: "wang",
		Age:  1,
		Male: true,
	}

	polink := jumper_transform.Newtransform()

	polink.AddOp(def.PacketJson, nil)
	//polink.AddOp(def.PacketXml, nil)

	polink.AddOp(def.CompressGzip, nil)
	//polink.AddOp(def.CompressZlib, nil)

	//polink.AddOp(def.EncryptAes, []interface{}{[]byte("abcdefghijklmnop")})
	//polink.AddOp(def.EncryptDes, []interface{}{[]byte("ijklmnop")})
	polink.AddOp(def.EncryptRsa, []interface{}{PubKeyRemote, PrivKeyLocal})
	var rst1 []byte
	err := polink.Execute(def.Forward, test1, &rst1)
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}

	var test2 s1
	err = polink.Execute(def.Backward, rst1, &test2)
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}

	fmt.Printf("result: %v\n", test2)
	fmt.Printf("origin data: %v\n", test1)
	if test1 != test2 {
		panic("err")
	}

	fmt.Println("--------------------------------------------------------")
	/////////////////////////// protobuf ///////////////////////////

	polink.Reset()
	//bodyData := "guangzhou/fangcun/vip/company"
	bodyData := "test protobuf"
	p := &StringMessage{
		Body: proto.String(bodyData),
		Header: &Header{
			MessageId: proto.String("20-05"),
			Topic:     proto.String("golang"),
		},
	}

	polink.AddOp(def.PacketProtobuf, nil)
	//polink.AddOp(def.CompressGzip, nil)
	//polink.AddOp(def.CompressZlib, nil)
	//polink.AddOp(def.EncryptAes, []interface{}{[]byte("abcdefghijklmnop")})
	//polink.AddOp(def.EncryptDes, []interface{}{[]byte("ijklmnop")})
	//polink.AddOp(def.EncryptRsa, []interface{}{PubKeyRemote, PrivKeyLocal})
	var rst11 []byte
	err = polink.Execute(def.Forward, p, &rst11)
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}

	fmt.Printf("=== rst: %x", rst11)
	var protobufrst StringMessage
	err = polink.Execute(def.Backward, rst11, &protobufrst)
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}

	fmt.Printf("result: %s\n", protobufrst.String())
	fmt.Printf("origin data: %s\n", p.String())
	if protobufrst.String() != p.String() {
		panic("err")
	}

	fmt.Println("--------------------------------------------------------")
	/////////////////////////// binary ///////////////////////////
	msg := &interf.Message{
		Type:    1,
		Content: []byte("this is a Message."),
	}
	polink.Reset()
	polink.AddOp(def.PacketBinary, nil)
	//polink.AddOp(def.CompressGzip, nil)
	polink.AddOp(def.CompressZlib, nil)
	//polink.AddOp(def.EncryptAes, []interface{}{[]byte("abcdefghijklmnop")})
	//polink.AddOp(def.EncryptDes, []interface{}{[]byte("ijklmnop")})
	polink.AddOp(def.EncryptRsa, []interface{}{PubKeyRemote, PrivKeyLocal})
	var rst12 []byte
	err = polink.Execute(def.Forward, msg, &rst12)
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}

	var msgrst interf.Message
	err = polink.Execute(def.Backward, rst12, &msgrst)
	if err != nil {
		fmt.Printf("err: %s", err)
		return
	}

	fmt.Printf("result: %v\n", msgrst)
	fmt.Printf("origin data: %v\n", *msg)
	if msgrst.Type != msg.Type || string(msgrst.Content) != string(msg.Content) {
		panic("err")
	}
}

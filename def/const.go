package def

//transform 表示一个操作，　包含三种类型：
//1. encode
//2. compress
//3. encrypt

const (
	PackageOpMin int8 = 0 + iota
	//封包
	PacketBase64
	PacketBase64RawUrl
	PacketJson
	PacketXml
	PacketProtobuf
	PacketBinary

	//压缩
	CompressGzip
	CompressZlib

	//加密
	EncryptMd5
	EncryptSha1
	EncryptSha256
	EncryptAes
	EncryptDes
	EncryptRsa

	PackageOpMax
)

const (
	Forward  int8 = 1 //打包->压缩->加密
	Backward int8 = 2 //解密->解压->解包
)

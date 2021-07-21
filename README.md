# jumper_transform
Encapsulate some common encoding, compression, encryption functions, provide abstract interfaces, and can be used in a convenient combination. 

封装一些常用的编码、压缩、加密函数，提供抽象接口，可以较方便的自由组合使用。

使用方法：
1. 调用 Newtransform 获取一个 interf.Transform 接口
2. 调用接口方法 AddOp 来添加想要进行的操作
3. 调用接口方法 Execute 来进行所添加的操作，具有方向性：\
   前向(Forward)：封包 -> 压缩 -> 加密\
   后向(Backward)：解密 -> 解压 -> 解包

可选操作如下：

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

实例参见目录 example
   
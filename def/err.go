package def

const (
	ErrMd5NoDecryptCode    = 11001
	ErrSha1NoDecryptCode   = 11002
	ErrSha256NoDecryptCode = 11003

	ErrInvalidAesKeyCode = 11004
	ErrInvalidDesKeyCode = 11005

	ErrInvalidRsaPrivateKeyCode = 11006
	ErrInvalidRsaPublicKeyCode  = 11007

	ErrParamShouldNotNilCode           = 11008
	ErrParamShouldStringOrBytesCode    = 11009
	ErrParamShouldPointOfByteSliceCode = 11010
	ErrParamShouldImplInterfMsgCode    = 11011
	ErrParamShouldImplProtoMsgCode     = 11012
	ErrDecodeFailedCode                = 11013
)

type Error struct {
	Code    int32
	Message string
}

func (this *Error) Error() string {
	return this.Message
}

func New(code int32, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

var (
	ErrMd5NoDecrypt                = New(ErrMd5NoDecryptCode, "md5 no decrypt.")
	ErrSha1NoDecrypt               = New(ErrSha1NoDecryptCode, "sha1 no decrypt.")
	ErrSha256NoDecrypt             = New(ErrSha256NoDecryptCode, "sha256 no decrypt.")
	ErrInvalidAesKey               = New(ErrInvalidAesKeyCode, "invalid aes key.")
	ErrInvalidDesKey               = New(ErrInvalidDesKeyCode, "invalid des key.")
	ErrInvalidRsaPrivateKey        = New(ErrInvalidRsaPrivateKeyCode, "invalid rsa private key.")
	ErrInvalidRsaPublicKey         = New(ErrInvalidRsaPublicKeyCode, "invalid rsa public key.")
	ErrParamShouldNotNil           = New(ErrParamShouldNotNilCode, "param should not be nil.")
	ErrParamShouldStringOrBytes    = New(ErrParamShouldStringOrBytesCode, "param should be string or []byte.")
	ErrParamShouldPointOfByteSlice = New(ErrParamShouldPointOfByteSliceCode, "param should be *[]byte.")
	ErrParamShouldImplInterfMsg    = New(ErrParamShouldImplInterfMsgCode, "param should be pointer of interf.Message.")
	ErrParamShouldImplProtoMsg     = New(ErrParamShouldImplProtoMsgCode, "param should implement proto.Message.")
	ErrDecodeFailed                = New(ErrDecodeFailedCode, "decode failed.")
)

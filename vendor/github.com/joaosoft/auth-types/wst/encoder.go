package wst

const (
	EncodeAscii85     encodeType = "ascii85"
	EncodeBase32      encodeType = "base32"
	EncodeBase64      encodeType = "base64"
	EncodeHexadecimal encodeType = "hexadecimal"
	EncodeCipher      encodeType = "cipher"
)

type encodeType string

type iencoder interface {
	Encode(value []byte) ([]byte, error)
	Decode(value []byte) ([]byte, error)
}

var cipherSecret []byte = []byte("a very very very very secret key")

func WithCipherSecret(secret []byte) {
	var secret32 [32]byte
	copy(secret32[:], secret)
	cipherSecret = secret32[:]
}

var encoderMethods = map[encodeType]iencoder{
	EncodeAscii85:     &encoderAscii85{},
	EncodeBase32:      &encoderBase32{},
	EncodeBase64:      &encoderBase64{},
	EncodeHexadecimal: &encoderHexadecimal{},
	EncodeCipher:      &encoderCipher{encoder: &encoderBase64{}, key: &cipherSecret},
}

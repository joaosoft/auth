package wst

import (
	"bytes"
	"encoding/base32"
)

type encoderBase32 struct{}

func (e *encoderBase32) Encode(value []byte) ([]byte, error) {
	base32Encoded := make([]byte, base32.StdEncoding.EncodedLen(len(value)))
	base32.StdEncoding.Encode(base32Encoded, value)

	return base32Encoded, nil
}

func (e *encoderBase32) Decode(value []byte) ([]byte, error) {
	base32Encoded := bytes.NewBuffer(value)
	base32Decoded := make([]byte, base32.StdEncoding.DecodedLen(len(base32Encoded.Bytes())))
	n, err := base32.StdEncoding.Decode(base32Decoded, base32Encoded.Bytes())

	return base32Decoded[0:n], err
}

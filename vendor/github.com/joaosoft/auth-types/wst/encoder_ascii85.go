package wst

import (
	"encoding/ascii85"
)

type encoderAscii85 struct{}

func (e *encoderAscii85) Encode(value []byte) ([]byte, error) {
	ascii85Encoded := make([]byte, ascii85.MaxEncodedLen(len(value)))
	n := ascii85.Encode(ascii85Encoded, value)

	return ascii85Encoded[0:n], nil
}

func (e *encoderAscii85) Decode(value []byte) ([]byte, error) {
	ascii85Decoded := make([]byte, ascii85.MaxEncodedLen(len(value)))
	n, _, err := ascii85.Decode(ascii85Decoded, value, true)

	return ascii85Decoded[0:n], err
}

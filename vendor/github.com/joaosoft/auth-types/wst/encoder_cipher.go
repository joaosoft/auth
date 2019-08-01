package wst

import (
	"crypto/aes"
	"crypto/cipher"
)

type encoderCipher struct {
	encoder iencoder
	key     *[]byte
}

func (e *encoderCipher) Encode(value []byte) ([]byte, error) {
	block, err := aes.NewCipher(*e.key)
	if err != nil {
		return nil, err
	}

	cipherText := make([]byte, aes.BlockSize+len(value))
	iv := cipherText[:aes.BlockSize]

	cfbEncripter := cipher.NewCFBEncrypter(block, iv)
	cfbEncripter.XORKeyStream(cipherText[aes.BlockSize:], value)

	return e.encoder.Encode(cipherText)
}

func (e *encoderCipher) Decode(value []byte) ([]byte, error) {
	text, err := e.encoder.Decode(value)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(*e.key)
	if err != nil {
		return nil, err
	}

	if len(text) < aes.BlockSize {
		return nil, ErrorCipherTextTooShort
	}

	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]

	cfbDecrypter := cipher.NewCFBDecrypter(block, iv)
	cfbDecrypter.XORKeyStream(text, text)

	return text, nil
}

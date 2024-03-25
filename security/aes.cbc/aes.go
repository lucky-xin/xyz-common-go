package aescbc

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

type Encryptor struct {
	Key string
	Iv  string
}

func (e *Encryptor) Encrypt(origData []byte) (r string, err error) {
	block, err := aes.NewCipher([]byte(e.Key))
	if err != nil {
		return
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, []byte(e.Iv))
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

func (e *Encryptor) Decrypt(crypted string) (re string, err error) {
	decodeData, err := base64.StdEncoding.DecodeString(crypted)
	if err != nil {
		return
	}
	block, err := aes.NewCipher([]byte(e.Key))
	if err != nil {
		return
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(e.Iv))
	origData := make([]byte, len(decodeData))
	blockMode.CryptBlocks(origData, decodeData)
	origData = PKCS5UnPadding(origData)
	return string(origData), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
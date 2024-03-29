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

func (e *Encryptor) Encrypt(origData []byte) (byts []byte, err error) {
	block, err := aes.NewCipher([]byte(e.Key))
	if err != nil {
		return
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, []byte(e.Iv))
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	byts = crypted
	return
}

func (e *Encryptor) EncryptWithBase64(origData []byte) (r string, err error) {
	byts, err := e.Encrypt(origData)
	if err != nil {
		return
	}
	return base64.StdEncoding.EncodeToString(byts), nil
}

func (e *Encryptor) Decrypt(decodeData []byte) (byts []byte, err error) {
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
	byts = PKCS5UnPadding(origData)
	return
}

func (e *Encryptor) DecryptBase64(crypted string) (byts []byte, err error) {
	decodeData, err := base64.StdEncoding.DecodeString(crypted)
	if err != nil {
		return
	}
	return e.Decrypt(decodeData)
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

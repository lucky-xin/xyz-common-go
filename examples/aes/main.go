package main

import "github.com/lucky-xin/xyz-common-go/security/aes.cbc"

func main() {
	key := "B31F2A75FBF94099B31F2A75FBF94099" // 此处16|24|32个字符
	iv := "1234567890123456"
	println(len(key), "--", len(iv))

	encryptor := aescbc.Encryptor{Key: key, Iv: iv}
	re, err := encryptor.Encrypt([]byte("Hello,World!"))
	if err != nil {
		panic(err)
	}
	println(string(re))
	byts, err := encryptor.Decrypt(re)
	if err != nil {
		panic(err)
	}
	println(string(byts))
	println(re)

	encryptor = aescbc.Encryptor{Key: key, Iv: iv}
	res, err := encryptor.EncryptWithBase64([]byte("Hello,World!"))
	if err != nil {
		panic(err)
	}
	println(res)
	byts, err = encryptor.DecryptBase64(res)
	if err != nil {
		panic(err)
	}
	println(string(byts))
	println(re)
}

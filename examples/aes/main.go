package main

import "github.com/lucky-xin/xyz-common-go/security/aes.cbc"

func main() {
	key := "B31F2A75FBF94099B31F2A75FBF94099" // 此处16|24|32个字符
	iv := "1234567890123456"
	encryptor := aescbc.Encryptor{Key: key, Iv: iv}
	re, err := encryptor.Decrypt("r4GYgavrMOzZRB8o+ctl1A==")
	if err != nil {
		panic(err)
	}
	println(encryptor.Encrypt([]byte("Hello,World!")))
	println(re)
}

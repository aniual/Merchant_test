package controllers

import (
	_"fmt"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"bytes"
)


var Key = []byte("B31F2A75FBF94099")

var Iv = []byte("1234567890123456")


/*func main() {

	s, _:= Decrypt("66p1ohg0GJ6rZNY24fUvYIOT2y/twIy+npYxccVV6rY=")
	fmt.Println(Encrypt([]byte(s)))
}*/


func Encrypt(origData []byte) (string, error) {
	block, err := aes.NewCipher(Key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, Iv)
	crypted := make([]byte, len(origData))

	blockMode.CryptBlocks(crypted, origData)
	return base64.StdEncoding.EncodeToString(crypted), nil
}


func Decrypt(crypted string) (string, error) {
	decodeData, err := base64.StdEncoding.DecodeString(crypted)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(Key)
	if err != nil {
		return "", err
	}
	//blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, Iv)
	origData := make([]byte, len(decodeData))
	blockMode.CryptBlocks(origData, decodeData)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return string(origData), nil
}


func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
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

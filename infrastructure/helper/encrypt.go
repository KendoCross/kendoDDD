package helper

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"errors"
	"fmt"
)

// md5加密 字符串
func Md5Encrypt(text string) string {
	data := []byte(text)
	has := md5.Sum(data)
	enText := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return enText
}

func EncryptPwd(text string) string {
	data := []byte(text + "nilinside")
	has := md5.Sum(data)
	enText := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return enText
}

func PF(v interface{}) {
	fmt.Printf("v : %+[1]v\n", v)
}

var ivAes = []byte("ba3b15e0e29211e9")

//加密
func AesCbcEncrypt(plainText, key []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("密钥长度必须为32位")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	paddingText := PKCS5Padding(plainText, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, ivAes)
	cipherText := make([]byte, len(paddingText))
	blockMode.CryptBlocks(cipherText, paddingText)
	return cipherText, nil
}

// decrypt
func AesCbcDecrypt(cipherText, key []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("密钥长度必须为32位")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, ivAes)
	paddingText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(paddingText, cipherText)
	plainText, err := PKCS5UnPadding(paddingText)
	if err != nil {
		return nil, err
	}
	return plainText, nil
}
func PKCS5Padding(plainText []byte, blockSize int) []byte {
	padding := blockSize - (len(plainText) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	newText := append(plainText, padText...)
	return newText
}
func PKCS5UnPadding(plainText []byte) ([]byte, error) {
	length := len(plainText)
	number := int(plainText[length-1])
	if number >= length {
		return nil, errors.New("填充尺寸错误")
	}
	return plainText[:length-number], nil
}

package cryptx

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
)

var sKey = "platshipzxcvbnmq"

func pkcs7Padding(data []byte, blockSize int) []byte {
	//判断缺少几位长度。最少1，最多 blockSize
	padding := blockSize - len(data)%blockSize
	//补足位数。把切片[]byte{byte(padding)}复制padding个
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// pkcs7UnPadding 填充的反向操作
func pkcs7UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("strings error")
	}
	//获取填充的个数
	unPadding := int(data[length-1])
	if unPadding > length {
		return nil, errors.New("strings error")
	}
	return data[:(length - unPadding)], nil
}

func validateLength(data []byte, blockSize int) error {
	if len(data)%blockSize != 0 {
		return fmt.Errorf("data length must be a multiple of block size (%d)", blockSize)
	}
	return nil
}

// AesEncrypt 加密
func AesEncrypt(str, iKey string) (d string, err error) {
	if str == "" {
		return "", errors.New("aes encrypted content cannot ne empty")
	}
	if iKey == "" {
		iKey = sKey
	}
	data := []byte(str)
	key := []byte(iKey)
	//创建加密实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	//判断加密块的大小
	blockSize := block.BlockSize()
	//填充
	encryptBytes := pkcs7Padding(data, blockSize)
	//初始化加密数据接收切片
	crypted := make([]byte, len(encryptBytes))
	//使用cbc加密模式
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	//执行加密
	blockMode.CryptBlocks(crypted, encryptBytes)
	return strings.Replace(base64.StdEncoding.EncodeToString(crypted), "+", "_", -1), nil
}

// AesDecrypt 解密
func AesDecrypt(str string, iKey string) (string, error) {
	if str == "" {
		return "", errors.New("aes decrypted content cannot be empty")
	}

	str = strings.Replace(str, "_", "+", -1)
	if iKey == "" {
		iKey = sKey
	}
	n, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}
	data := []byte(n)
	key := []byte(iKey)
	//创建实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	//获取块的大小
	blockSize := block.BlockSize()
	// 校验数据长度
	if err := validateLength(data, blockSize); err != nil {
		return "", err
	}
	//使用cbc
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	//初始化解密数据接收切片
	crypted := make([]byte, len(data))
	//执行解密
	blockMode.CryptBlocks(crypted, data)
	//去除填充
	crypted, err = pkcs7UnPadding(crypted)
	if err != nil {
		return "", err
	}
	return string(crypted), nil
}

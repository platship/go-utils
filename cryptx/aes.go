/*
 * @Author: coller
 * @Date: 2023-02-08 10:04:16
 * @LastEditors: coller
 * @LastEditTime: 2024-04-21 14:36:38
 * @Desc:
 */
package cryptx

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

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
		return nil, errors.New("加密字符串错误！")
	}
	//获取填充的个数
	unPadding := int(data[length-1])
	return data[:(length - unPadding)], nil
}

// AesEncrypt 加密
func AesEncrypt(s, sKey, iKey string) (d string) {
	data := []byte(s)
	key := []byte(sKey)
	if iKey != "" {
		key = []byte(iKey)
	}
	//创建加密实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return ""
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
	return base64.StdEncoding.EncodeToString(crypted)
}

// AesDecrypt 解密
func AesDecrypt(s, sKey, iKey string) string {
	n, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return ""
	}
	data := []byte(n)
	key := []byte(sKey)
	if iKey != "" {
		key = []byte(iKey)
	}
	//创建实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return ""
	}
	//获取块的大小
	blockSize := block.BlockSize()
	//使用cbc
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	//初始化解密数据接收切片
	crypted := make([]byte, len(data))
	//执行解密
	blockMode.CryptBlocks(crypted, data)
	//去除填充
	crypted, err = pkcs7UnPadding(crypted)
	if err != nil {
		return ""
	}
	return string(crypted)
}

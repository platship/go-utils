/*
 * @Author: Coller
 * @Date: 2021-09-24 12:30:08
 * @LastEditTime: 2024-04-21 14:41:39
 * @Desc: hash获取
 */
package cryptx

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"

	"github.com/fasthey/go-utils/conv"
)

/**
 * @desc: MD5哈希值
 * @param undefined
 * @return {*}
 */
func MD5(b []byte) string {
	h := md5.New()
	_, _ = h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// MD5String MD5哈希值
func MD5String(s string) string {
	return MD5(conv.StringToByte(s))
}

// SHA1 SHA1哈希值
func SHA1(b []byte) string {
	h := sha1.New()
	_, _ = h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// SHA1String SHA1哈希值
func SHA1String(s string) string {
	return SHA1(conv.StringToByte(s))
}

/*
 * @Author: Coller
 * @Date: 2021-09-24 12:30:08
 * @LastEditTime: 2024-04-21 14:44:52
 * @Desc: md5
 */
package cryptx

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/fasthey/go-utils/conv"
)

/**
 * @desc: md5 encryption
 * @return {*}
 */
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write(conv.StringToByte(value))

	return hex.EncodeToString(m.Sum(nil))
}

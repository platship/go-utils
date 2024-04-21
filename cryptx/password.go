/*
 * @Author: Coller
 * @Date: 2021-09-24 12:30:08
 * @LastEditTime: 2024-04-21 14:37:44
 * @Desc: 密码验证
 */
package cryptx

import (
	"fmt"

	"github.com/fasthey/go-utils/conv"

	"golang.org/x/crypto/bcrypt"
)

/**
 * @desc: 密码加密
 * @param 加密的字符串
 * @return {*}
 */
func EncodePassword(rawPassword string) string {
	hash, err := bcrypt.GenerateFromPassword(conv.StringToByte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	return conv.ByteToString(hash)
}

/**
 * @desc: 密码验证
 * @param 验证的密码
 * @param 输入的密码
 * @return {*}
 */
func ValidatePassword(encodePassword, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword(conv.StringToByte(encodePassword), conv.StringToByte(inputPassword))
	return err == nil
}

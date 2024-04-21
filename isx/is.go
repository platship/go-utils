/*
 * @Author: Coller
 * @Date: 2022-05-17 12:21:52
 * @LastEditTime: 2024-04-21 14:42:29
 * @Desc: 数据判断
 */
package isx

import (
	"errors"
	"regexp"
	"strings"
	"unicode"

	"github.com/fasthey/go-utils/stringx"
)

/**
 * @desc: 是否为字符串
 * @param undefined
 * @return {*}
 */
func IsString(str string) bool {
	strLen := len(str)
	if str == "" || strLen == 0 {
		return true
	}
	for i := 0; i < strLen; i++ {
		if unicode.IsSpace(rune(str[i])) == false {
			return false
		}
	}
	return true
}

/**
 * @desc: 是否包含字符串
 * @param undefined
 * @return {*}
 */
func IsAnyString(strs ...string) bool {
	for _, str := range strs {
		if IsString(str) {
			return true
		}
	}
	return false
}

/**
 * @desc: 是否为空
 * @param undefined
 * @return {*}
 */
func IsEmpty(str string) bool {
	return len(str) == 0
}

/**
 * @desc: 字符串是否相等
 * @param undefined
 * @param undefined
 * @return {*}
 */
func IsEquals(a, b string) bool {
	return a == b || strings.ToUpper(a) == strings.ToUpper(b)
}

/**
 * @desc: IsUsername 验证用户名合法性，用户名必须由5-12位(数字、字母、_、-)组成，且必须以字母开头。
 * @param undefined
 * @return {*}
 */
func IsUsername(username string) error {
	if IsString(username) {
		return errors.New("请输入用户名")
	}
	matched, err := regexp.MatchString("^[0-9a-zA-Z_-]{5,12}$", username)
	if err != nil || !matched {
		return errors.New("用户名必须由5-12位(数字、字母、_、-)组成，且必须以字母开头")
	}
	matched, err = regexp.MatchString("^[a-zA-Z]", username)
	if err != nil || !matched {
		return errors.New("用户名必须由5-12位(数字、字母、_、-)组成，且必须以字母开头")
	}
	return nil
}

/**
 * @desc: IsEmail 验证是否是合法的邮箱
 * @param undefined
 * @return {*}
 */
func IsEmailErr(email string) (err error) {
	if !IsString(email) {
		err = errors.New("邮箱格式不符合规范")
		return
	}
	pattern := `^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,4})$`
	matched, _ := regexp.MatchString(pattern, email)
	if !matched {
		err = errors.New("邮箱格式不符合规范")
	}
	return nil
}

func IsEmail(email string) bool {
	if err := IsEmailErr(email); err != nil {
		return false
	}
	return true
}

/**
 * @desc: 验证字段是否合法
 * @param 允许 a-z小写字母 4-30位
 * @return {*}
 */
func IsField(name string) (err error) {
	if IsString(name) {
		return errors.New("字段格式不正确")
	}
	matched, err := regexp.MatchString("^[a-z_]{4,30}$", name)
	if err != nil || !matched {
		return errors.New("字段标识必须由4-30位小写字母或下划线组成组成，且必须以字母开头")
	}
	firstChar := string([]rune(name)[:1])
	if firstChar == "_" {
		return errors.New("不能以下划线开头")
	}
	return nil
}

/**
 * @desc: IsPassword 是否是合法的密码
 * @param undefined
 * @param undefined
 * @return {*}
 */
func IsPassword(password, rePassword string) error {
	if IsString(password) {
		return errors.New("请输入密码")
	}
	if stringx.RuneLen(password) < 6 {
		return errors.New("密码过于简单")
	}
	if password != rePassword {
		return errors.New("两次输入密码不匹配")
	}
	return nil
}

/**
 * @desc: IsURL 是否是合法的URL
 * @param undefined
 * @return {*}
 */
func IsURL(url string) bool {
	if IsString(url) {
		return false
	}
	indexOfHttp := strings.Index(url, "http://")
	indexOfHttps := strings.Index(url, "https://")
	if indexOfHttp == 0 || indexOfHttps == 0 {
		return true
	}
	return false
}

/**
 * @desc: 是否是合法手机号
 * @param undefined
 * @return {*}
 */
func IsMobile(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

/**
 * @desc: 字符串列表是否有重复数据
 * @param {[]string} list
 * @return {*}
 */
func IsStringListDup(list []string) bool {
	tmpMap := make(map[string]int)
	for _, value := range list {
		tmpMap[value] = 1
	}
	var keys []interface{}
	for k := range tmpMap {
		keys = append(keys, k)
	}
	if len(keys) != len(list) {
		return true
	}
	return false
}

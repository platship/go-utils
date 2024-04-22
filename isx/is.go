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
)

/**
 * @desc: IsUsername 验证用户名合法性，用户名必须由5-12位(数字、字母、_、-)组成，且必须以字母开头。
 * @param undefined
 * @return {*}
 */
func IsUsername(username string) error {
	if username == "" {
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
 * @desc: 验证字段是否合法
 * @param 允许 a-z小写字母 4-30位
 * @return {*}
 */
func IsField(name string) (err error) {
	if name == "" {
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

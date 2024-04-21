/*
 * @Author: coller
 * @Date: 2024-04-10 09:55:10
 * @LastEditors: coller
 * @LastEditTime: 2024-04-10 09:56:42
 * @Desc:
 */
package numberx

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

/**
 * @desc: uint数组去重
 * @param {[]uint} list
 * @return {*}
 */
func UintArrayRemoveDup(list []uint) []uint {
	var x []uint = []uint{}
	for _, i := range list {
		if len(x) == 0 {
			x = append(x, i)
		} else {
			for k, v := range x {
				if i == v {
					break
				}
				if k == len(x)-1 {
					x = append(x, i)
				}
			}
		}
	}
	return x
}

/**
 * @desc: 查询uint在不在数组中
 * @param ids 查询的数组
 * @param id 要查询的值
 * @return {*}
 */
func UintInArray(id uint, ids []uint) bool {
	for _, eachId := range ids {
		if eachId == id {
			return true
		}
	}
	return false
}

func UintArrayRemove(s []uint, index uint) []uint {
	return append(s[:index], s[index+1:]...)
}

/**
 * @desc:生成随机数字
 * @param start 开始
 * @param end 结束
 * @return {*}
 */
func RandInt(start int, end int) string {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(end - start)
	random = start + random
	return strconv.Itoa(random)
}

func MoneyToFormatFloat(old string) float64 {
	s := strings.Replace(old, "$", "", 1)
	s = strings.Replace(s, ",", "", -1)
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

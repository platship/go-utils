package numberx

import (
	"crypto/rand"
	"math/big"
	mathRand "math/rand"
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
	mathRand.Seed(time.Now().UnixNano())
	random := mathRand.Intn(end - start)
	random = start + random
	return strconv.Itoa(random)
}

// RandRangeInt 获取范围随机数 [min, max)
func RandRangeInt[T int | int64](min, max T) T {
	if min < 0 || max <= 0 {
		return 0
	}
	if min >= max {
		return 0
	}
	maxBigInt := big.NewInt(int64(max))
	i, _ := rand.Int(rand.Reader, maxBigInt)
	if i.Int64() < int64(min) {
		return RandRangeInt(min, max)
	}
	return T(i.Int64())
}

func MoneyToFormatFloat(old string) float64 {
	s := strings.Replace(old, "$", "", 1)
	s = strings.Replace(s, ",", "", -1)
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

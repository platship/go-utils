/*
 * @Author: Coller
 * @Date: 2022-05-17 12:38:10
 * @LastEditTime: 2024-04-21 17:17:40
 * @Desc: 字符串处理
 */
package stringx

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/fasthey/go-utils/conv"

	"github.com/fasthey/go-utils/phpx"

	"golang.org/x/crypto/scrypt"
)

/**
 * @desc: 截取字符串
 * @param s 字符串
 * @param start 开始的位置
 * @param length 长度
 * @return {*}
 */
func CutString(s string, start, length int) string {
	bt := []rune(s)
	if start < 0 {
		start = 0
	}
	if start > len(bt) {
		start = start % len(bt)
	}
	var end int
	if (start + length) > (len(bt) - 1) {
		end = len(bt)
	} else {
		end = start + length
	}
	return string(bt[start:end])
}

/**
 * @desc: 随机获取字符串
 * @param l 长度
 * @return {*}
 */
func RandString(l int, types ...string) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	num := "0123456789"
	sym := "_-"
	var res string
	if len(types) > 0 {
		if types[0] == "string" {
			res = str
		} else if types[0] == "xid" {
			res = num + "abcdefghijklmnopqrstuvwxyz"
		} else if types[0] == "number" {
			res = num
		}
	} else {
		res = str + num + sym
	}
	bytes := conv.StringToByte(res)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return conv.ByteToString(result)
}

/**
 * @desc: 根据明文密码和加盐值生成密码
 * @param password
 * @param salt 盐值
 * @return {*}
 */
func GetPassword(password string, salt string) (verify string, err error) {
	var rb []byte
	rb, err = scrypt.Key(conv.StringToByte(password), conv.StringToByte(salt), 16384, 8, 1, 32)
	if err != nil {
		return
	}
	verify = hex.EncodeToString(rb)
	return
}

// 去除前后无用字符
func StringTrim(str string, characterMask ...string) string {
	if len(characterMask) == 0 {
		return strings.TrimSpace(str)
	}
	return strings.Trim(str, characterMask[0])
}

/**
 *  @Description: 字符串补零
 *  @param str :需要操作的字符串
 *  @param resultLen 结果字符串的长度
 *  @param reverse true 为前置补零，false 为后置补零
 *  @return string
 */
func ZeroFillByStr(str string, resultLen int, reverse bool) string {
	if len(str) > resultLen || resultLen <= 0 {
		return str
	}
	if reverse {
		return fmt.Sprintf("%0*s", resultLen, str) // 不足前置补零
	}
	result := str
	for i := 0; i < resultLen-len(str); i++ {
		result += "0"
	}
	return result
}

func ListDup(list []string) []string {
	dupFre := make(map[string]int)
	var dep []string
	for _, item := range list {
		// 检查重复频率map中是否存在项目/元素
		_, exist := dupFre[item]
		if exist {
			//如果已经在map中，则将计数器增加1
			dupFre[item] += 1
			dep = append(dep, item)
		} else {
			//从1开始计数
			dupFre[item] = 1
		}
	}
	return dep
}

func FormatMoney(money float64) string {
	return phpx.NumberFormat(money, 2, ".", "")
}

// UpperFirst converts the first character of string to upper case.
func UpperFirst(s string) string {
	if len(s) == 0 {
		return ""
	}

	r, size := utf8.DecodeRuneInString(s)
	r = unicode.ToUpper(r)

	return string(r) + s[size:]
}

// LowerFirst converts the first character of string to lower case
func LowerFirst(s string) string {
	if len(s) == 0 {
		return ""
	}

	r, size := utf8.DecodeRuneInString(s)
	r = unicode.ToLower(r)

	return string(r) + s[size:]
}

/**
 * @desc: RuneLen 字符成长度
 * @param undefined
 * @return {*}
 */
func RuneLen(s string) int {
	bt := []rune(s)
	return len(bt)
}

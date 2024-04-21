/*
 * @Author: Coller
 * @Date: 2022-01-20 12:07:15
 * @LastEditTime: 2024-02-21 10:47:32
 * @Desc: 数据转换
 */
package conv

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unsafe"
)

/**
 * @desc: 字节转字符串
 * @return {*}
 */
func ByteToString(b []byte) string {
	/* #nosec G103 */
	return *(*string)(unsafe.Pointer(&b))
}

/**
 * @desc: 字符串转字节
 * @return {*}
 */
func StringToByte(s string) (b []byte) {
	/* #nosec G103 */
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	/* #nosec G103 */
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	return b
}

/**
 * @desc: int转字符串
 * @param undefined
 * @return {*}
 */
func IntToString(i int) string {
	return strconv.Itoa(i)
}

/**
 * @desc: uint转字符串
 * @param undefined
 * @return {*}
 */
func UintToString(i uint) string {
	return strconv.Itoa(int(i))
}

/**
 * @desc: 字符串转int
 * @param undefined
 * @return {*}
 */
func StringToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return num
}

/**
 * @desc: 字符串转Uint
 * @param undefined
 * @return {*}
 */
func StringToUint(i string) uint {
	return uint(StringToInt(i))
}

/**
 * @desc: 字符串转float
 * @param undefined
 * @return {*}
 */
func StringToFloat64(s string) float64 {
	float, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return float
}

/**
 * @desc: 字符串转uint数组
 * @param 传入的字符串
 * @param 分隔符，如 “,”
 * @return {*}
 */
func StringToIntArray(s, sep string) []int {
	strArr := strings.Split(s, sep)
	var n []int
	for _, v := range strArr {
		n = append(n, StringToInt(v))
	}
	return n
}

/**
 * @desc: 字符串转uint数组
 * @param 传入的字符串
 * @param 分隔符，如 “,”
 * @return {*}
 */
func StringToUintArray(s, sep string) []uint {
	s = strings.Trim(s, sep)
	strArr := strings.Split(s, sep)
	var n []uint
	for _, v := range strArr {
		n = append(n, uint(StringToInt(v)))
	}
	return n
}

/**
 * @desc: uint数组转string
 * @param undefined
 * @return {*}
 */
func UintArrayToString(a []uint) string {
	return strings.Replace(strings.Trim(fmt.Sprint(a), "[]"), " ", ",", -1)
}

/**
 * @desc: 字符数组转string
 * @param undefined
 * @return {*}
 */
func StringArrayToString(a []string) string {
	return strings.Replace(strings.Trim(fmt.Sprint(a), "[]"), " ", ",", -1)
}

/**
* @desc: 下划线转小驼峰
 */
func NameToCamel(s string) string {
	if strings.HasSuffix(s, "_cate_id") {
		return strings.Replace(s, "_cate_id", "CateName", 1)
	} else if strings.HasSuffix(s, "_user_id") {
		return strings.Replace(s, "_user_id", "UserName", 1)
	} else if strings.HasSuffix(s, "_dept_id") {
		return strings.Replace(s, "_dept_id", "DeptName", 1)
	} else if strings.HasSuffix(s, "_id") {
		return strings.Replace(s, "_id", "Name", 1)
	} else {
		return SnakeToCamel(s)
	}

}

func NameToSnake(s string) string {
	if strings.HasSuffix(s, "CateName") {
		return strings.Replace(s, "CateName", "_cate_id", 1)
	} else if strings.HasSuffix(s, "UserName") {
		return strings.Replace(s, "UserName", "_user_id", 1)
	} else if strings.HasSuffix(s, "DeptName") {
		return strings.Replace(s, "DeptName", "_dept_id", 1)
	} else if strings.HasSuffix(s, "Name") {
		return strings.Replace(s, "Name", "_id", 1)
	} else {
		return CamelToSnake(s)
	}
}

/**
 * 蛇形转驼峰
 * @description xx_yy to XxYx  xx_y_y to XxYY
 * @date 2020/7/30
 * @param s要转换的字符串
 * @return string
 **/
func SnakeToCamel(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	s = string(data[:])
	return strings.ToLower(s[:1]) + s[1:]
}

/**
 * @desc: 小写驼峰转下划线
 * @param {string} s
 * @return {*}
 */
func CamelToSnake(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
		} else {
			if unicode.IsUpper(r) {
				output = append(output, '_')
			}
			output = append(output, unicode.ToLower(r))
		}
	}
	return string(output)
}

/**
 * @desc: 浮点数转字符串，用于金额
 * @param f 浮点数
 * @param m 尾部位数
 * @return {*}
 */
func FloatToString(f float64, m int) string {
	n := strconv.FormatFloat(f, 'f', -1, 32)
	if n == "" {
		return ""
	}
	if m >= len(n) {
		return n
	}
	newNum := strings.Split(n, ".")
	if len(newNum) < 2 || m >= len(newNum[1]) {
		if m == 2 { // 如果是金额
			return n + ".00"
		} else {
			return n
		}
	}
	return newNum[0] + "." + newNum[1][:m]
}

/**
 * @desc: 浮点数转字符串，用于金额
 * @param f 浮点数
 * @param m 尾部位数
 * @return {*}
 */
func InterfaceToUintArr(val interface{}) (res []uint, err error) {
	value := reflect.ValueOf(val)
	if value.Kind() != reflect.Slice && value.Kind() != reflect.Array {
		return nil, errors.New("parse error")
	}
	for i := 0; i < value.Len(); i++ {
		res = append(res, uint(11))
	}
	return res, nil
}

func InterfaceToString(value interface{}) string {
	var key string
	if value == nil {
		return key
	}
	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case time.Time:
		t, _ := value.(time.Time)
		key = t.String()
		// 2022-11-23 11:29:07 +0800 CST  这类格式把尾巴去掉
		key = strings.Replace(key, " +0800 CST", "", 1)
		key = strings.Replace(key, " +0000 UTC", "", 1)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}
	return key
}

func InterfaceToInt(value interface{}) int {
	var key int
	if value == nil {
		return key
	}
	switch value.(type) {
	case float64, float32:
		ft := value.(float64)
		key = int(ft)
	case int:
		it := value.(int)
		key = it
	case uint:
		it := value.(uint)
		key = int(it)
	case int8:
		it := value.(int8)
		key = int(it)
	case int16:
		it := value.(int16)
		key = int(it)
	case uint8:
		it := value.(uint8)
		key = int(it)
	case uint16:
		it := value.(uint16)
		key = int(it)
	case int32:
		it := value.(int32)
		key = int(it)
	case int64:
		it := value.(int64)
		key = int(it)
	case uint32:
		it := value.(uint32)
		key = int(it)
	case uint64:
		it := value.(uint64)
		key = int(it)
	default:
		key = 0
	}
	return key
}

/**
 * @desc: 字节转浮点
 * @param undefined
 * @return {*}
 */
func ByteToFloat(b []byte) float64 {
	v2, _ := strconv.ParseFloat(ByteToString(b), 32/64)
	return v2
}

/**
 * @desc: 字节转int
 * @param undefined
 * @return {*}
 */
func Byte2Int(b []byte) int {
	return StringToInt(ByteToString(b))
}

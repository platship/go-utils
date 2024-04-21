/*
 * @Author: Coller
 * @Date: 2022-05-17 21:31:09
 * @LastEditTime: 2024-04-10 12:25:01
 * @Desc: 数组处理
 */
package stringx

/**
 * @desc: 查询字符串在不在数组里面
 * @param items 查询的数组
 * @param item 查询的字符串
 * @return {*}
 */
func StringInArray(item string, items []string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

/**
 * @desc: 字符串数组去重
 * @param {[]uint} list
 * @return {*}
 */
func StringArrayRemoveDup(list []string) []string {
	var x []string = []string{}
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

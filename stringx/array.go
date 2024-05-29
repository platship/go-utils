/*
 * @Author: Coller
 * @Date: 2022-05-17 21:31:09
 * @LastEditTime: 2024-04-21 17:35:29
 * @Desc: 数组处理
 */
package stringx

func ArrayRemoveElement(arr []string, target string) []string {
	var newArray []string
	for _, num := range arr {
		if num != target {
			newArray = append(newArray, num)
		}
	}
	return newArray
}

/**
 * @desc: 字符串数组去重
 * @param {[]uint} list
 * @return {*}
 */
func ArrayRemoveDup(list []string) []string {
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

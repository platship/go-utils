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

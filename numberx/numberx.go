package numberx

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

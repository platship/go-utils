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

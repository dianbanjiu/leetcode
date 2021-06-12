package perday

// leetcode-cn number: 5767. 检查是否区域内所有整数都被覆盖
func IsCovered(ranges [][]int, left int, right int) bool {
	var changed bool
	for _, r := range ranges {
		if r[0] <= left && r[1] >= left {
			left = r[1]
			changed = true
		}	
		if r[0] <= right && r[1] >= right {
			right = r[0]
			changed = true
		}

		if right <= left && changed{
			return true
		}
	}
	
	return false
}
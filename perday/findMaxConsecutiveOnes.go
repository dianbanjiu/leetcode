package perday

// leetcode-cn number 485
// url: https://leetcode-cn.com/problems/max-consecutive-ones/
// 方法1：计算连续1最大的数
func FindMaxConsecutiveOnes1(nums []int) int {
	var onesnum, onestemp int
	for _, num := range nums {
		if onestemp == onestemp+num {
			if onesnum < onestemp {
				onesnum = onestemp
			}
			onestemp = 0
			continue
		}
		onestemp += num
	}

	if onesnum < onestemp {
		onesnum = onestemp
	}
	return onesnum
}
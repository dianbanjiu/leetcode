package src

import "sort"

// 判断一副牌是否是顺子
func IsStraight(nums []int) bool {
	sort.Ints(nums)
	var zeroCounts int
	if nums[0] == 0 {
		zeroCounts++
	}

	var difference int
	for i := 1; i < 5; i++ {
		if nums[i] == 0 {
			zeroCounts++
		} else if nums[i-1] == nums[i] {
			return false
		} else if nums[i-1] != 0 {
			difference += nums[i] - nums[i-1] - 1
		}
	}

	if difference-zeroCounts <= 0 {
		return true
	}
	return false
}

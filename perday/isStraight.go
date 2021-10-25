package perday

import "sort"

// 剑指 Offer 61. 扑克牌中的顺子
func isStraight(nums []int) bool {
	sort.Ints(nums)
	var joker int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			joker++
			continue
		}
		if nums[i] == nums[i+1] {
			return false
		}
	}
	if nums[len(nums)-1]-nums[joker] < 5 {
		return true
	}
	return false
}

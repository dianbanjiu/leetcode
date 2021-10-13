package perday

import "sort"

// 剑指 Offer 53 - I. 在排序数组中查找数字 I
func search(nums []int, target int) int {
	leftIndex := sort.SearchInts(nums, target)
	if leftIndex == len(nums) {
		return 0
	}

	rightIndex := sort.SearchInts(nums, target+1) - 1
	return rightIndex - leftIndex + 1
}

package perday

import "sort"

// 剑指 Offer 39. 数组中出现次数超过一半的数字
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

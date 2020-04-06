package src

import "sort"

func FindKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	k = len(nums) - k%len(nums)
	if k == len(nums) {
		return nums[0]
	}
	return nums[k]
}

package src

import (
	"sort"
)

func TopKFrequent(nums []int, k int) []int {
	var r = make([]int, k)
	sort.Ints(nums)
	r[0] = nums[0]
	j := 1
	for i := 1; i < len(nums) && j < k; i++ {
		if r[j-1] != nums[i] {
			r[j] = nums[i]
			j++
		}
	}
	return r
}

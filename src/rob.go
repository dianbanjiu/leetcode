package src

import "math"

func Rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var r = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			r[0] = nums[0]
		} else if i == 1 {
			r[1] = int(math.Max(float64(nums[1]), float64(nums[0])))
		} else {
			r[i] = int(math.Max(float64(r[i-2]+nums[i]), float64(r[i-1])))
		}
	}
	return r[len(nums)-1]
}

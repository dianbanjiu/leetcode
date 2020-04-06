package src

import "sort"

func LengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var dp = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		var max int
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] && dp[j]+1 > max {
				max = dp[j] + 1
			}
		}
		dp[i] = max
	}
	sort.Ints(dp)
	return dp[len(nums)-1] + 1
}

package src

func MinStartValue(nums []int) int {
	var dp = make([]int, len(nums))
	dp[0] = nums[0]
	min := dp[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = dp[i-1] + nums[i]
		if dp[i] < min {
			min = dp[i]
		}
	}
	if min < 1 {
		return 1 - min
	}
	return 1
}

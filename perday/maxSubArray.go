package perday

// 剑指 Offer 42. 连续子数组的最大和
func maxSubArray(nums []int) int {
	var max = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] = nums[i-1] + nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return nums[len(nums)-1]
}

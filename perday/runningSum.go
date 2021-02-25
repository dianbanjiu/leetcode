package perday

// leetcode-cn number: 1480. 一维数组的动态和
// url:https://leetcode-cn.com/problems/running-sum-of-1d-array/
func RunningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}

	return nums
}

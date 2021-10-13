package perday

// leetcode-cn number: 5800. 基于排列构建数组
func BuildArray(nums []int) []int {
	var ans = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = nums[nums[i]]
	}
	return ans
}

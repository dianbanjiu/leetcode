package perday

// leetcode-cn number: 5808. 数组串联
func getConcatenation(nums []int) []int {
	var numsLen = len(nums)
	var ans = make([]int, numsLen*2)

	for i := 0; i < numsLen; i++ {
		ans[i] = nums[i]
		ans[i+numsLen] = nums[i]
	}

	return ans
}

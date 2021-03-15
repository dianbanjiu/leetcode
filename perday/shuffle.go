package perday

// leetcode-cn number: 1470. 重新排列数组
// url: https://leetcode-cn.com/problems/shuffle-the-array/
func Shuffle(nums []int, n int) []int {
	var result = make([]int, len(nums))
	var i, j, k = 0, n, 0
	for i<n && j < n*2 {
		result[k], result[k+1] = nums[i], nums[j]
		k+=2
		i++
		j++
	}

	return result
}
package perday

// leetcode-cn number: 503. 下一个更大元素 II
// url: https://leetcode-cn.com/problems/next-greater-element-ii/
func NextGreaterElements(nums []int) []int {
	var result = make([]int, len(nums))
	for i, num := range nums {
		for j := (i + 1) % len(nums); j < len(nums); j = (j + 1) % len(nums) {
			if num < nums[j] {
				result[i] = nums[j]
				break
			}

			if i == j {
				result[i] = -1
				break
			}
		}
	}
	return result
}

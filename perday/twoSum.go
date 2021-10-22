package perday

// 剑指 Offer 57. 和为s的两个数字
func twoSum(nums []int, target int) []int {
	p, l := 0, len(nums)-1
	for p < l {
		sum := nums[p] + nums[l]
		if sum == target {
			return []int{nums[p], nums[l]}
		}
		if sum > target {
			l--
		} else {
			p++
		}
	}
	return nil
}

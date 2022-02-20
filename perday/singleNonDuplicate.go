package perday

// leetcode-cn nunber: 540. 有序数组中的单一元素
// leetcode-cn url: https://leetcode-cn.com/problems/single-element-in-a-sorted-array/
func singleNonDuplicate(nums []int) int {
	var result = nums[len(nums)-1]
	for i := 0; i < len(nums)-1; i += 2 {
		if nums[i] == nums[i+1] {
			continue
		}
		result = nums[i]
		break
	}
	return result
}

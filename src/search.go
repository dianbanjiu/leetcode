package src

// 计算数组中给定元素出现的次数
// 计算方法：二分查找
func Search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	var r int
	if nums[len(nums)/2] == target {
		r += 1
		r += Search(nums[:len(nums)/2], target)
		r += Search(nums[len(nums)/2+1:], target)
	} else if nums[len(nums)/2] < target {
		r += Search(nums[len(nums)/2+1:], target)
	} else if nums[len(nums)/2] > target {
		r += Search(nums[:len(nums)/2], target)
	}
	return r
}

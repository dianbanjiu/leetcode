package perday

// leetcode-cn number: 35. 搜索插入位置
// leetcode-cn url: https://leetcode-cn.com/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	index := binarySearch(nums, 0, len(nums)-1, target)
	return index
}

func binarySearch(nums []int, pre, last, target int) int {
	if pre > last {
		return pre
	}
	middle := (pre + last) / 2
	if nums[middle] == target {
		return middle
	} else if nums[middle] > target {
		return binarySearch(nums, pre, middle-1, target)
	} else {
		return binarySearch(nums, middle+1, last, target)
	}
}

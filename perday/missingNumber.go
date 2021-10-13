package perday

// 剑指 Offer 53 - II. 0～n-1中缺失的数字
func missingNumber(nums []int, target int) int {
	pre, last := 0, len(nums)-1
	for pre <= last {
		mid := (pre + last) / 2
		if mid == nums[mid] {
			pre = mid + 1
		} else {
			last = mid - 1
		}

	}
	return pre
}

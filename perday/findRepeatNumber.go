package perday

func FindRepeatNumber(nums []int) int {
	for i := 0; ; {
		if i != nums[i] && nums[i] == nums[nums[i]] {
			return nums[i]
		}
		if i != nums[i] {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		} else {
			i++
		}
	}
}

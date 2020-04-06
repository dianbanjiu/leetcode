package src

func FindPeakElement(nums []int) int {
	if nums[0] > nums[1] {
		return nums[0]
	} else if nums[len(nums)-1] > nums[len(nums)-2] {
		return nums[len(nums)-1]
	}
	var r int
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			r = i
			break
		}
	}
	return r
}

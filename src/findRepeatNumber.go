package src

func findRepeatNumber(nums []int) int {
	var cases = make(map[int]int)
	var r int
	for i := 0; i < len(nums); i++ {
		if _, ok := cases[nums[i]]; ok {
			r = nums[i]
			break
		} else {
			cases[nums[i]] = 1
		}
	}
	return r
}

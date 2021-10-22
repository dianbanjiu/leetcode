package perday

// 剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
func Exchange(nums []int) []int {
	p, l := 0, len(nums)-1
	for p < l {
		py := nums[p] % 2
		ly := nums[l] % 2
		if py == 1 {
			p++
		} else if py == 0 && ly == 1 {
			nums[p], nums[l] = nums[l], nums[p]
			p++
			l--
		} else {
			l--
		}
	}
	return nums
}

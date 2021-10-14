package perday

// 剑指 Offer 11. 旋转数组的最小数字
func MinArray(numbers []int) int {
	p, l := 0, len(numbers)-1
	for p < l {
		mid := (p + l) / 2

		if numbers[mid] > numbers[l] {
			p = mid + 1
		} else if numbers[mid] < numbers[l] {
			l = mid
		} else {
			l -= 1
		}

	}
	return numbers[p]
}

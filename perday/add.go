package perday

// 剑指 Offer 65. 不用加减乘除做加法
func add(a int, b int) int {

	for b != 0 {
		c := (a & b) << 1
		a ^= b
		b = c
	}
	return a
}

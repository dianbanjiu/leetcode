package perday

// 剑指 Offer 66. 构建乘积数组
func ConstructArr(a []int) []int {
	if len(a) == 0 {
		return a
	}
	var b = make([]int, len(a))
	b[0] = 1
	for i := 1; i < len(a); i++ {
		b[i] = b[i-1] * a[i-1]
	}

	tmp := a[len(a)-1]
	for i := len(a) - 2; i >= 0; i-- {
		b[i] = b[i] * tmp
		tmp *= a[i]
	}

	return b
}

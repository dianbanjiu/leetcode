package src

// 构建乘积数组
func ConstructArr(a []int) []int {
	if len(a) == 0 {
		return nil
	}
	var fdp = make([]int, len(a))
	var edp = make([]int, len(a))
	fdp[0] = 1
	edp[len(a)-1] = 1
	for i := 1; i < len(a); i++ {
		fdp[i] = a[i-1] * fdp[i-1]
	}

	for i := len(a) - 2; i >= 0; i-- {
		edp[i] = a[i+1] * edp[i+1]
	}

	var ans = make([]int, len(a))
	for i := 0; i < len(a); i++ {
		ans[i] = fdp[i] * edp[i]
	}
	return ans
}

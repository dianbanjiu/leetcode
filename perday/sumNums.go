package perday

// 剑指 Offer 64. 求1+2+…+n
func sumNums(n int) int {
	var res int
	sum(&res, n)
	return res
}

func sum(res *int, n int) bool {
	*res += n
	return n > 0 && sum(res, n-1)
}

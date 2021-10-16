package perday

// 剑指 Offer 10- II. 青蛙跳台阶问题
func numWays(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	var result = [2]int{1, 2}
	for i := 3; i < n; i++ {
		result[0], result[1] = result[1], (result[0]+result[1])%1000000007
	}

	return (result[0] + result[1]) % 1000000007
}

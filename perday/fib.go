package perday

// 剑指 Offer 10- I. 斐波那契数列
func fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	var fibResult = [2]int{0, 1}
	for i := 2; i < n; i++ {
		fibResult[0], fibResult[1] = fibResult[1]%1000000007, (fibResult[0]+fibResult[1])%1000000007
	}
	return (fibResult[0] + fibResult[1]) % 1000000007
}

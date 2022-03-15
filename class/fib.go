package class

// 斐波那契数列
func Fib(n int) int{
	if n == 0 {
		return 0
	}else if n == 1{
		return 1
	}
	num1, num2 := 0, 1
	var result int
	for i := 2; i <= n; i++ {
		result = num1+num2 
		num1, num2 = num2, result
	}
	return result
}
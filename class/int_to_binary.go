package class

// 整数转二进制
func IntToBinary(n int) string {
	if n == 0 {
		return "0"
	}
	var result string
	for n != 0 {
		temp := n
		if n>>1<<1 != temp {
			result = "1" + result
		} else {
			result = "0" + result
		}
		n >>= 1
	}
	return result
}

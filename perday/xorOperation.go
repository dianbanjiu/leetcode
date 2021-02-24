package perday

// leetcode-cn number:1486. 数组异或操作
// url: https://leetcode-cn.com/problems/xor-operation-in-an-array/
func XorOperation(n int, start int) int {
	var result int
	for i := 0; i < n; i++ {
		result ^= start + 2*i
	}
	return result
}

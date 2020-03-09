package src

// 解题的思想类似于整数的反转
// 将返回结果左移一位，取出原数的最后一位跟返回结果做和
// 重复上面一步直到原数的位数全部用完
func ReverseBits(num uint32) uint32 {
	var r uint32
	for i := 0; i < 32; i++ {
		r = (r << 1) | (num & 1)
		num >>= 1
	}
	return r
}

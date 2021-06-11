package perday

// leetcode-cn number: 461. 汉明距离
// leetcode-cn url: https://leetcode-cn.com/problems/hamming-distance/
func hammingDistance(x int, y int) int {
	xorValue := x ^ y
	var count int
	for xorValue != 0 {
		if (xorValue>>1)<<1 != xorValue {
			count++
		}
		xorValue = xorValue >> 1
	}
	return count
}

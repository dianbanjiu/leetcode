package perday

// leetcode-cn number: 剑指 Offer 15. 二进制中1的个数
// url: https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/
func hammingWeight(num uint32) int {
	var count int
	for num != 0 {
		if (num>>1)<<1 != num {
			count++
		}
		num>>=1
	}
	return count
}
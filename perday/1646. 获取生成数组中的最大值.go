package perday

// leetcode-cn number: 1646. 获取生成数组中的最大值
// leetcode-cn url: https://leetcode-cn.com/problems/get-maximum-in-generated-array/
// tag: 动态规划
func getMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	var result = make([]int, n+1)
	result[0], result[1] = 0, 1
	var max = 1
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			result[i] = result[i/2]
		} else {
			result[i] = result[i/2] + result[i/2+1]
		}
		if result[i] > max {
			max = result[i]
		}
	}
	return max
}

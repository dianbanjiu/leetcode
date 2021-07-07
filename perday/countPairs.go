package perday

/*
	leetcode-cn number: 1711. 大餐计数
	url: https://leetcode-cn.com/problems/count-good-meals/
 	解题思路：
			找到数组中最大的数 maxValue，则数组中任意两数的和都小于等于 maxValue*2
		创建一个 map 存储所有元素出现的次数，计算数组中任意两个元素和小于等于 maxValue*2 的 2 的幂
		因为题目设置 maxValue*2 的最大值为 2^20，也就意味着每个元素在这个范围内最多需要将 1 左移 20 次
*/
func countPairs(deliciousness []int) int {
	var maxValue = deliciousness[0]
	for _, de := range deliciousness {
		if de > maxValue {
			maxValue = de
		}
	}
	var allElements = make(map[int]int)
	allElements[deliciousness[0]] = 1
	var count int
	var maxSum = maxValue * 2
	for i := 1; i < len(deliciousness); i++ {
		for sum := 1; sum <= maxSum; sum <<= 1 {
			if allElements[sum-deliciousness[i]] == 0 {
				continue
			}
			count += allElements[sum-deliciousness[i]]
		}
		allElements[deliciousness[i]] += 1

	}
	return count % 1000000007
}

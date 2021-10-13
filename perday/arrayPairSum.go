package perday

import "sort"

// leetcode-cn number 561. 数组拆分 I
// url: https://leetcode-cn.com/problems/array-partition-i/
// 解题思路：根据题意可以总结一个规律，就是所求的 min(ai, bi) 的总和其实就是将数组数据由小到达排序，
// 将所有数据分为 n 组，取每组第一个数
func ArrayPairSum(nums []int) int {
	var sum int
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}

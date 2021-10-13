package perday

import "sort"

//leetcode-cn number: 1833. 雪糕的最大数量
//leetcode url: https://leetcode-cn.com/problems/maximum-ice-cream-bars/
func MaxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	var sum, count int
	for _, cost := range costs {
		sum += cost
		if sum > coins {
			break
		}
		count++
	}
	return count
}

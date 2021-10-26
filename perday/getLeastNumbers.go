package perday

import "sort"

// 剑指 Offer 40. 最小的k个数
func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}

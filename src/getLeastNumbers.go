package src

import "sort"

func GetLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	for i := 1; i < len(arr); {
		if arr[i] == arr[i-1] {
			arr = append(arr[:i-1], arr[i:]...)
			continue
		}
		i++
	}
	return arr[:k]
}

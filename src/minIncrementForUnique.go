package src

import (
	"sort"
)

func MinIncrementForUnique(A []int) int {
	sort.Ints(A)
	var count int
	for i := 1; i < len(A); i++ {
		if A[i] <= A[i-1] {
			count += A[i-1] - A[i] + 1
			A[i] += A[i-1] - A[i] + 1
		}
	}
	return count
}

package src

import (
	"math"
)

func SumFourDivisors(nums []int) int {
	var sum int
	for i := 0; i < len(nums); i++ {
		t := div(nums[i])
		if len(t) == 4 {
			sum += t[0] + t[1] + t[2] + t[3]
		}
	}
	return sum
}

func div(n int) []int {
	var t = make([]int, 0)
	for j := 1; j <= int(math.Sqrt(float64(n))); j++ {
		if n%j == 0 && n/j != j {
			t = append(t, j, n/j)
		} else if n%j == 0 && n/j == j {
			t = append(t, j)
		}
	}
	return t
}

// 1,2,3,4,5,6,7,8,9,10
// 6:1,2,3,6=12
// 8:1,2,4,8=15
// 10:1,2,5,10=18

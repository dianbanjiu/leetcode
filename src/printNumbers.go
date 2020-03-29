package src

import (
	"math"
)

func PrintNumbers(n int) []int {
	ansLen := int(math.Pow10(n))-1
	var ans = make([]int, ansLen)
	for i:=0; i<ansLen; i++{
		ans[i]=i+1
	}
	return ans
}
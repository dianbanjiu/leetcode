package leetcode

import (
	"leetcode/src"
	"testing"
)

func TestSumFourDivisors(t *testing.T) {
	var (
		in       = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		excepted = 32
	)
	actual := src.SumFourDivisors(in)
	if actual != excepted {
		t.Error("excepted ", excepted, "but got ", actual)
	}
}

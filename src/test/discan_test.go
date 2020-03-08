package leetcode

import (
	"leetcode/src"
	"testing"
)

func TestDistributeCandies(t *testing.T) {
	var demo = []struct {
		c      int
		n      int
		except []int
	}{
		{7, 4, []int{1, 2, 3, 1}},
		{10, 3, []int{5, 2, 3}},
	}
	for _, v := range demo {
		actual := src.DistributeCandies(v.c, v.n)
		for i, _ := range actual {
			if actual[i] != v.except[i] {
				t.Error(actual, v.except)
				break
			}
		}
	}
}

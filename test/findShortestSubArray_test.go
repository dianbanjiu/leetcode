package leetcodeTest

import (
	"leetcode/perday"
	"testing"
)

func TestFindShortestSubArray(t *testing.T) {
	var cases = []struct {
		in     []int
		except int
	}{
		{[]int{1, 2, 2, 3, 1}, 2},
		{[]int{12}, 1},
		{[]int{1, 2, 2, 3, 1, 4, 2}, 6},
		{[]int{1, 1}, 2},
		{[]int{1, 1, 2, 2, 2, 1}, 3},
		{[]int{1, 2, 1}, 3},
	}

	for _, s := range cases {
		output := perday.FindShortestSubArray(s.in)
		if s.except != output {
			t.Errorf("in %v, except %v, but got %v", s.in, s.except, output)
		}
	}
}

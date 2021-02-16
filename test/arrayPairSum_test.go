package leetcodeTest

import (
	"leetcode/perday"
	"testing"
)

func TestArrayPairSum(t *testing.T) {
	var cases = []struct {
		In []int
		Except int
	}{
		{[]int{1,4,3,2}, 4},
		{[]int{6,2,6,5,1,2}, 9},
	}

	for _, s := range cases {
		out := perday.ArrayPairSum(s.In)
		if out != s.Except {
			t.Errorf("in %v, except %v, but got %v", s.In, s.Except, out)
		}
	}
}
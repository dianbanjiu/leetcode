package leetcodeTest

import (
	"leetcode/perday"
	"reflect"
	"testing"
)

func TestCountBits(t *testing.T) {
	var cases = []struct {
		num    int
		expect []int
	}{
		{2, []int{0, 1, 1}},
		{5, []int{0, 1, 1, 2, 1, 2}},
		{100, []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5, 2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6, 2, 3, 3, 4, 3}},
	}
	for _, s := range cases {
		output := perday.CountBits(s.num)
		if !reflect.DeepEqual(output, s.expect) {
			t.Errorf("in %v, expect %v, but got %v", s.num, s.expect, output)
		}
	}
}

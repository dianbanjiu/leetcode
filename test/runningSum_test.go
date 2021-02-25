package leetcodeTest

import (
	"leetcode/perday"
	"reflect"
	"testing"
)

func TestRunningSum(t *testing.T) {
	var cases = []struct {
		nums   []int
		expect []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}},
	}

	for _, s := range cases {
		output := perday.RunningSum(s.nums)
		if !reflect.DeepEqual(output, s.expect) {
			t.Errorf("in %v, expect %v, but got %v", s.nums, s.expect, output)
		}
	}
}

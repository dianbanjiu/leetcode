package leetcodeTest

import (
	"leetcode/perday"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	var caseTest = []struct {
		in  []int
		out bool
	}{
		{[]int{}, false},
		{[]int{1}, false},
		{[]int{1, 1}, true},
		{[]int{1, 2}, false},
		{[]int{1, 2, 1}, true},
		{[]int{1, 1, 2}, true},
	}

	for _, value := range caseTest {
		actual := perday.ContainsDuplicate(value.in)
		if actual != value.out {
			t.Errorf("in %v, except %v, actual %v", value.in, actual, value.out)
		}
	}
}

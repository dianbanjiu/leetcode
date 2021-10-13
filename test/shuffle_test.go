package leetcodeTest

import (
	"leetcode/perday"
	"reflect"
	"testing"
)

func TestShuffle(t *testing.T) {
	var demos = []struct {
		nums   []int
		n      int
		except []int
	}{
		{[]int{2, 5, 1, 3, 4, 7}, 3, []int{2, 3, 5, 4, 1, 7}},
		{[]int{1, 2, 3, 4, 4, 3, 2, 1}, 4, []int{1, 4, 2, 3, 3, 2, 4, 1}},
		{[]int{1, 1, 2, 2}, 2, []int{1, 2, 1, 2}},
	}

	for _, demo := range demos {
		output := perday.Shuffle(demo.nums, demo.n)
		if !reflect.DeepEqual(output, demo.except) {
			t.Errorf("nums is %v, n is %v, except %v, but got %v", demo.nums, demo.n, demo.except, output)
		}
	}
}

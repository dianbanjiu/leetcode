package leetcodeTest

import (
	"leetcode/perday"
	"reflect"
	"testing"
)

func TestNextGreaterElements(t *testing.T) {
	var demos = []struct {
		nums   []int
		except []int
	}{
		{[]int{1, 2, 1}, []int{2, -1, 2}},
		{[]int{1, 2, 3, 4, 3}, []int{2, 3, 4, -1, 4}},
	}

	for _, demo := range demos {
		output := perday.NextGreaterElements(demo.nums)
		if !reflect.DeepEqual(output, demo.except) {
			t.Errorf("nums is %v, except %v, but got %v", demo.nums, demo.except, output)
		}
	}
}

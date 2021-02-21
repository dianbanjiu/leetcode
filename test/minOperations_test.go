package leetcodeTest

import (
	"leetcode/perday"
	"reflect"
	"testing"
)

func TestMinOperations(t *testing.T) {
	var cases = []struct{
		boxes string
		except []int
	}{
		{"110", []int{1,1,3}},
		{"001011", []int{11,8,5,4,3,4}},
	}

	for _, s := range cases {
		output := perday.MinOperations(s.boxes)
		if !reflect.DeepEqual(output, s.except) {
			t.Errorf("in %v, except %v, but got %v", s.boxes,s.except, output)
		}
	}
}
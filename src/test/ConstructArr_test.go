package leetcode

import (
	"leetcode/src"
	"reflect"
	"testing"
)

func TestConstructArr(t *testing.T) {
	var testCases = []struct {
		in     []int
		except []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{120, 60, 40, 30, 24}},
		{[]int{}, nil},
	}

	for _, tc := range testCases {
		actual := src.ConstructArr(tc.in)
		if !reflect.DeepEqual(actual, tc.except) {
			t.Errorf("excepted %v, got %v", tc.except, actual)
		}
	}
}

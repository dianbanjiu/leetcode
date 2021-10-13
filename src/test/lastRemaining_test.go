package leetcode

import (
	"leetcode/src"
	"testing"
)

func TestLastRemaining(t *testing.T) {
	var testCases = []struct {
		n      int
		m      int
		except int
	}{
		{5, 3, 3},
		{10, 17, 2},
		{2, 2, 0},
	}

	for _, tc := range testCases {
		actual := src.LastRemaining(tc.n, tc.m)
		if actual != tc.except {
			t.Errorf("except %v got %v", tc.except, actual)
		}
	}
}

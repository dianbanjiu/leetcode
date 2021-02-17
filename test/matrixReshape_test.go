package leetcodeTest

import (
	"leetcode/perday"
	"reflect"
	"testing"
)

func TestMatrixReshape(t *testing.T)  {
	var cases = []struct{
		nums [][]int
		r int
		c int
		except [][]int
	}{
		{[][]int{{1,2},{3,4}}, 1,4,[][]int{{1,2,3,4}}},
		{[][]int{{1,2},{3,4}},2,4, [][]int{{1,2},{3,4}}},
	}

	for _, s := range cases {
		out := perday.MatrixReshape(s.nums, s.r, s.c)
		if !reflect.DeepEqual(out, s.except) {
			t.Error("out:", out, "except: ", s.except)
		}
	}
}

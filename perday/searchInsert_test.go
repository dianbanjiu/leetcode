package perday

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	var testCase = []struct {
		inNum    []int
		inTarget int
		Except   int
	}{
		{[]int{1, 3, 5, 6}, 5, 2},
		{[]int{1, 3, 5, 6}, 2, 1},
		{[]int{1, 3, 5, 6}, 0, 0},
		{[]int{1}, 0, 0},
		{[]int{1, 3, 5, 6}, 7, 4},
	}

	ast := assert.New(t)
	for _, tC := range testCase {
		ast.Equal(tC.Except, searchInsert(tC.inNum, tC.inTarget))
	}
}

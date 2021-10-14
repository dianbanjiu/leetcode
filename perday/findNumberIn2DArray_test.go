package perday

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindNumberIn2DArray(t *testing.T) {
	var testCases = []struct {
		inMatrix [][]int
		inTarget int
		out      bool
	}{
		{
			[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}},
			5, true,
		}, {
			[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}},
			20, false,
		}, {
			[][]int{{-5}}, -5, true,
		},
	}

	ast := assert.New(t)

	for _, tC := range testCases {
		ast.Equal(FindNumberIn2DArray(tC.inMatrix, tC.inTarget), tC.out)
	}
}

package perday

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxValue(t *testing.T) {
	testCases := []struct {
		in  [][]int
		out int
	}{
		{
			[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 12,
		}, {
			[][]int{}, 0,
		},
	}
	ast := assert.New(t)
	for _, tC := range testCases {
		ast.Equal(tC.out, MaxValue(tC.in))
	}
}

package perday

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinArray(t *testing.T) {
	testCases := []struct {
		in  []int
		out int
	}{
		{
			[]int{3, 4, 5, 1, 2}, 1,
		}, {
			[]int{2, 2, 2, 0, 1}, 0,
		}, {
			[]int{1, 3, 5}, 1,
		}, {
			[]int{1, 3}, 1,
		}, {
			[]int{1, 3, 3}, 1,
		},
	}

	ast := assert.New(t)
	for _, tC := range testCases {
		ast.Equal(tC.out, MinArray(tC.in))
	}
}

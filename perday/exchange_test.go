package perday

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExchange(t *testing.T) {
	var testCase = []struct {
		in  []int
		out []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 3, 2, 4}},
		{[]int{2, 16, 3, 5, 13, 1, 16, 1, 12, 18, 11, 8, 11, 11, 5, 1}, []int{3, 5, 13, 1, 1, 11, 11, 11, 5, 1, 2, 16, 16, 12, 18, 8}},
	}

	ast := assert.New(t)
	for _, tC := range testCase {
		ast.Equal(tC.out, Exchange(tC.in))
	}
}

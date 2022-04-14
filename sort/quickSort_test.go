package sortCode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuick(t *testing.T) {
	testCases := []struct {
		in     []int
		except []int
	}{
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{}, []int{}},
		{[]int{8, 1, 7, 5, 3, 6, 4, 2, 9}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	ast := assert.New(t)
	for _, tC := range testCases {
		Quick(tC.in)
		ast.ElementsMatch(tC.except, tC.in)
	}
}

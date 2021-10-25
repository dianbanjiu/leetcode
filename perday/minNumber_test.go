package perday

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinNumber(t *testing.T) {
	var testCases = []struct {
		in  []int
		out string
	}{
		{[]int{10, 2}, "102"},
		{[]int{3, 30, 34, 5, 9}, "3033459"},
	}

	ast := assert.New(t)
	for _, testCase := range testCases {
		ast.Equal(testCase.out, minNumber(testCase.in))
	}
}

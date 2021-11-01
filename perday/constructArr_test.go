package perday

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConstructArr(t *testing.T) {
	var testCases = []struct {
		in  []int
		out []int
	}{
		{[]int{1, 2, 3, 4, 5},
			[]int{120, 60, 40, 30, 24}},
	}

	ast := assert.New(t)
	for _, testCase := range testCases {
		ast.Equal(testCase.out, ConstructArr(testCase.in))
	}
}

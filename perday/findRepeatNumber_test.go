package perday

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 剑指 Offer 03. 数组中重复的数字
func TestFindRepeatNumber(t *testing.T) {

	var testCase = []struct {
		in  []int
		out int
	}{
		{[]int{2, 3, 1, 0, 2, 5, 3}, 2},
	}

	ast := assert.New(t)
	for _, tC := range testCase {
		ast.Equal(FindRepeatNumber(tC.in), tC.out)
	}
}

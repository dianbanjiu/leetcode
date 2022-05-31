package perday

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var testCase = []struct {
		nums   []int
		target int
		except []int
	}{
		{[]int{2, 7, 11, 5}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{}, 8, nil},
	}

	req := require.New(t)
	for _, tC := range testCase {
		req.ElementsMatch(tC.except, TwoSum(tC.nums, tC.target))
	}
}

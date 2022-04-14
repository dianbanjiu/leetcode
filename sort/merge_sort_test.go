package sortCode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMergeSort(t *testing.T) {
	var testCase = []struct {
		in     []int
		except []int
		desc   string
	}{
		{[]int{}, []int{}, "empty array "},
		{nil, nil, "nil array"},
		{[]int{1, 8, 6, 4, 5, 2, 3}, []int{1, 2, 3, 4, 5, 6, 8}, "normal array"},
	}

	var reqir = require.New(t)
	for _, tC := range testCase {
		reqir.ElementsMatch(tC.except, mergeSort(tC.in), tC.desc)
	}
}

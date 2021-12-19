package sortCode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertion(t *testing.T) {
	testCases := []struct {
		in     []int
		except []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2, 3}, []int{2, 3}},
		{[]int{1, 4, 6, 2, 3, 8, 5, 7, 9, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}
	td := assert.New(t)
	for _, tC := range testCases {
		Insertion(tC.in)
		td.Equal(tC.except, tC.in)
	}
}

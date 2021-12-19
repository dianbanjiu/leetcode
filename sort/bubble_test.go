package sortCode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubble(t *testing.T) {
	var cases = []struct {
		in  []int
		out []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2, 3}, []int{2, 3}},
		{[]int{1, 4, 6, 2, 3, 8, 5, 7, 9, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	td := assert.New(t)
	for _, cs := range cases {
		Bubble(cs.in)
		td.Equal(cs.out, cs.in)
	}
}

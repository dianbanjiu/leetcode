package perday

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := []struct {
		in  string
		out int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"abba", 2},
	}

	ast := assert.New(t)
	for _, tC := range testCases {
		ast.Equal(tC.out, LengthOfLongestSubstring(tC.in))
	}
}

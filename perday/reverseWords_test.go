package perday

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	var testCases = []struct {
		in  string
		out string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world!  ", "world! hello"},
		{"", ""},
		{" ", ""},
	}

	ast := assert.New(t)
	for _, testCase := range testCases {
		ast.Equal(testCase.out, ReverseWords(testCase.in))
	}
}

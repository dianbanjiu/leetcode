package perday

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstUniqChar(t *testing.T) {
	var testCase = []struct {
		in  string
		out byte
	}{
		{"abaccdeff", 'b'},
		{"bbcdea", 'c'},
	}

	ast := assert.New(t)
	for _, tC := range testCase {
		ast.Equal(FirstUniqChar(tC.in), tC.out)
	}
}

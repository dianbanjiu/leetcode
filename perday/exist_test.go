package perday

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExist(t *testing.T) {
	var testCase = []struct {
		inBoard [][]byte
		inWord  string
		out     bool
	}{
		{[][]byte{{'a'}}, "a", true},
	}

	ast := assert.New(t)
	for _, tc := range testCase {
		ast.Equal(tc.out, exist(tc.inBoard, tc.inWord))
	}
}

package perday

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseStr(t *testing.T) {
	testCases := []struct {
		inS    string
		inK    int
		except string
	}{
		{"abcdefg", 2, "bacdfeg"},
		{"abcd", 2, "bacd"},
		{"abcksjdlnfjddefg", 3, "cbaksjnldfjdfedg"},
	}

	ast := assert.New(t)
	for _, tC := range testCases {
		ast.Equal(tC.except, ReverseStr(tC.inS, tC.inK))
	}
}

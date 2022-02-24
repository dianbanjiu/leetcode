package perday

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	var testCase = []struct {
		in     string
		except []string
	}{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"", []string{}},
		{"2", []string{"a", "b", "c"}},
	}

	ast := assert.New(t)
	for _, tC := range testCase {
		ast.ElementsMatch(tC.except, letterCombinations(tC.in))
	}
}

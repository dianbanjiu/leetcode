package class

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
// 0 1 1 2 3 5 8 13
// 0 1 2 3 4 5 6 7
func TestFib(t *testing.T) {
	var testCase = []struct {
		in     int
		except int
	}{
		{0, 0},
		{1, 1},
		{5, 5},
		{7, 13},
	}
	ast := assert.New(t)
	for _, tC := range testCase {
		ast.Equal(tC.except, Fib(tC.in))
	}
}

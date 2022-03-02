package perday

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMaximumGenerated(t *testing.T) {
	var testCase = []struct {
		in     int
		except int
	}{
		{7, 3},
		{2, 1},
		{3, 2},
	}

	ast := assert.New(t)
	for _, tC := range testCase {
		ast.Equal(tC.except, getMaximumGenerated(tC.in))
	}
}

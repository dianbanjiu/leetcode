package perday

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertToTitle(t *testing.T) {
	var testCase = []struct {
		In     int
		Except string
	}{
		{1, "A"},
		{701, "ZY"},
		{28, "AB"},
		{2147483647, "FXSHRXW"},
		{702, "ZZ"},
	}

	ast := assert.New(t)
	for _, tC := range testCase {
		ast.Equal(tC.Except, convertToTitle(tC.In))
	}
}

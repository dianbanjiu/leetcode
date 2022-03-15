package class

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToBinary(t *testing.T){
	var testCase = []struct{
		in int
		except string
	}{
		{0,"0"},
		{1, "1"},
		{10, "1010"},
		{16, "10000"},
	}

	ast := assert.New(t)
	for _, tC := range testCase {
		ast.Equal(tC.except, IntToBinary(tC.in))
	}
}
package class

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYanghui(t *testing.T) {
	var testCase = []struct{
		in int
		except [][]int
	}{
		{1, [][]int{{1}}},
		{4, [][]int{{1},{1,1},{1,2,1}, {1,3,3,1}}},
		{6, [][]int{{1},{1,1},{1,2,1}, {1,3,3,1}, {1,4,6,4,1}, {1,5,10,10,5,1}}},
	}

	ast := assert.New(t)
	for _, tC := range testCase {
		ast.ElementsMatch(tC.except, Yanghui(tC.in))
	}
}

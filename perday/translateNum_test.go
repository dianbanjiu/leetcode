package perday

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslateNum(t *testing.T) {
	testCases := []struct {
		in  int
		out int
	}{
		{1, 1},
		{12, 2},
		{122, 3},
		{1225, 5},
		{12258, 5},
	}
	ast := assert.New(t)
	for _, tC := range testCases {
		ast.Equal(TranslateNum(tC.in), tC.out)
	}
}

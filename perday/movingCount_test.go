package perday

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMovingCount(t *testing.T) {
	var testCases = []struct {
		inM int
		inN int
		inK int
		out int
	}{
		{7, 2, 3, 7},
		{2, 3, 1, 3},
		{3, 1, 0, 1},
		{16, 8, 4, 15},
		{3, 2, 17, 6},
	}
	ast := assert.New(t)
	for _, testCase := range testCases {
		ast.Equal(testCase.out, movingCount(testCase.inM, testCase.inN, testCase.inK))
	}
}
func TestSumNumber(t *testing.T) {
	var testCases = []struct {
		in  int
		out int
	}{
		{1, 1},
		{12, 3},
		{12121, 7},
	}
	ast := assert.New(t)
	for _, testCase := range testCases {
		ast.Equal(testCase.out, sumNumber(testCase.in))
	}
}

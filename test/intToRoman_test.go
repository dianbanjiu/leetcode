package leetcodeTest

import (
	"leetcode/perday"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	var cases = []struct {
		in     int
		except string
	}{
		{1, "I"},
		{3999, "MMMCMXCIX"},
		{3, "III"},
		{4, "IV"},
		{9, "IX"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	}

	for _, s := range cases {
		output := perday.IntToRoman(s.in)
		if output != s.except {
			t.Errorf("in %v, except %v, got %v", s.in, s.except, output)
		}
	}
}

package leetcodeTest

import (
	"leetcode/perday"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	var cases = []struct{
		in string
		except int
	}{
		{"42", 42},
		{"-42", -42},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648},
	}

	for _, s := range cases {
		output := perday.MyAtoi(s.in)
		if output != s.except {
			t.Errorf("in %v, except %v, but got %v", s.in, s.except, output)
		}
	}
}

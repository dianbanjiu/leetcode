package leetcode

import (
	"leetcode/src"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	var demo = []struct {
		in     string
		except int
	}{
		{"V", 5},
		{"IV", 4},
	}
	for _, v := range demo {
		actual := src.RomanToInt(v.in)
		if actual != v.except {
			t.Errorf("input %s, except %d, but got %d", v.in, v.except, actual)
		}

	}
}

package leetcodeTest

import (
	"leetcode/perday"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	var demos = []struct {
		S      string
		except string
	}{
		{"abbaca", "ca"},
	}

	for _, demo := range demos {
		output := perday.RemoveDuplicates(demo.S)
		if output != demo.except {
			t.Errorf("S is %v, except is %v, but got %v", demo.S, demo.except, output)
		}
	}
}

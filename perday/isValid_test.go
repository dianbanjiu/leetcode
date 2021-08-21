package perday

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	var cases = []struct {
		in     string
		except bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}

	for _, cs := range cases {
		actual := IsValid(cs.in)
		if actual != cs.except {
			t.Errorf("in %s, except %v, but got %v. ", cs.in, cs.except, actual)
		}
	}
}

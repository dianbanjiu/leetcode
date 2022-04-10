package perday

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUniqueMorseRepresentations(t *testing.T) {
	var testCase = []struct {
		in     []string
		except int
		desc   string
	}{
		{[]string{"gin", "zen", "gig", "msg"}, 2, "符合规则的字符串"},
		{[]string{"a"}, 1, "临界字符串"},
	}

	var requir = require.New(t)
	for _, tC := range testCase {
		requir.Equal(tC.except, uniqueMorseRepresentations(tC.in))
	}
}

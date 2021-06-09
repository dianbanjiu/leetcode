package perday

import "testing"

func TestCountVowelStrings(t *testing.T){
	var demos = []struct{
		in int
		except int
	}{
		{1, 5},
		{2, 15},
		{3, 35},
		{4, 70},
		{5, 126},
		{33,66045},
	}

	for _, demo := range demos {
		actual := CountVowelStrings(demo.in)
		if actual!=demo.except {
			t.Errorf("in %d, except %d, but got %d", demo.in, demo.except, actual)
		}
	}
}
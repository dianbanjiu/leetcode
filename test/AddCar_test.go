package leetcodeTest

import (
	"leetcode/perday"
	"testing"
)

func TestAddCar(t *testing.T) {
	var parkingSystem = perday.CarConstructor(1, 1, 0)
	var cases = []struct {
		in     int
		expect bool
	}{
		{1, true},
		{2, true},
		{3, false},
		{1, false},
	}

	for _, s := range cases {
		var real = parkingSystem.AddCar(s.in)
		if real != s.expect {
			t.Errorf("in %v, expect %v, but got %v", s.in, s.expect, real)
		}
	}
}

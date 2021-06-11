package perday

import "testing"

func TestMaxSatisfaction(t *testing.T){
	var demos = []struct{
		in []int
		except int
	}{
		{[]int{-1,-8,0,5,-9},14},
		{[]int{4,3,2},20},
		{[]int{-1,-4,-5}, 0},
		{[]int{-2,5,-1,0,3,-3},35},
	}

	for _, demo := range demos {
		actual := MaxSatisfaction(demo.in)
		if actual!=demo.except {
			t.Errorf("in %v, except %d, but got %d", demo.in, demo.except, actual)
		}
	}
}
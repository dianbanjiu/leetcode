package perday

import (
	"reflect"
	"testing"
)

func TestCountPoints(t *testing.T){
	var demos = []struct{
		points [][]int
		queries [][]int
		except []int
	}{
		{[][]int{{1,3},{3,3},{5,3},{2,2}}, [][]int{{2,3,1},{4,3,1},{1,1,2}},[]int{3,2,2}},
		{[][]int{{1,1},{2,2},{3,3},{4,4},{5,5}},[][]int{{1,2,2},{2,2,2},{4,3,2},{4,3,3}},[]int{2,3,2,4}},
	}

	for _,demo := range demos {
		actual := CountPoints(demo.points,demo.queries)
		if !reflect.DeepEqual(actual, demo.except) {
			t.Errorf("points is %v, queries is %v, except %v, but got %v", demo.points,demo.queries,demo.except,actual)
		}
	}
}
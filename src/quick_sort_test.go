package src

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	testCases := []struct {
		desc string
		data []int
		want []int
	}{
		{
			"empty test",
			[]int{}, []int{},
		},
		{
			"test",
			[]int{1, 3, 4, 6, 2},
			[]int{1, 2, 3, 4, 6},
		},
		{
			"test1",
			[]int{3, 2, 7, 1, 6, 8, 5, 9, 4},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			"test2",
			[]int{34, 28, 27, 22, 14, 12, 8, 3, 0},
			[]int{0, 3, 8, 12, 14, 22, 27, 28, 34},
		},
		{
			"test3",
			[]int{6, 10, 11, 11, 18, 20, 23, 31, 34},
			[]int{6, 10, 11, 11, 18, 20, 23, 31, 34},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			Sort(tC.data)
			if !reflect.DeepEqual(tC.want, tC.data) {
				t.Errorf("want %v but got %v", tC.want, tC.data)
			}
		})
	}
}

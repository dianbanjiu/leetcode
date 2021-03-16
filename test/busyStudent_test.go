package leetcodeTest

import (
	"leetcode/perday"
	"testing"
)

func TestBusyStudent(t *testing.T) {
	var demos = []struct {
		startTime []int
		endTime   []int
		queryTime int
		except    int
	}{
		{[]int{1, 2, 3}, []int{3, 2, 7}, 4, 1},
		{[]int{4}, []int{4}, 4, 1},
		{[]int{1, 1, 1, 1}, []int{1, 3, 2, 4}, 7, 0},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{10, 10, 10, 10, 10, 10, 10, 10, 10}, 5, 5},
	}

	for _, demo := range demos {
		output := perday.BusyStudent(demo.startTime, demo.endTime, demo.queryTime)
		if output != demo.except {
			t.Errorf("startTime is %v, endTime is %v, queryTime is %v, except %v, but got %v", demo.startTime, demo.endTime, demo.queryTime, demo.except, output)
		}
	}
}

package src

import "reflect"

func PrisonAfterNDays(cells []int, N int) []int {
	if N == 0 {
		return cells
	}
	var cases = make(map[int][]int)
	for i := 0; i < N; i++ {
		t := cells[0]
		for j := 1; j < 7; j++ {
			if t == cells[j+1] {
				cells[j], t = 1, cells[j]
			} else {
				cells[j], t = 0, cells[j]
			}
		}

		cells[0] = 0
		cells[7] = 0
		s := append([]int{}, cells...)
		if reflect.DeepEqual(s, cases[0]) {
			break
		}
		cases[i] = s
	}
	N = (N - 1) % len(cases)
	return cases[N]
}

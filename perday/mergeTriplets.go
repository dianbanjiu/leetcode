package perday

import "reflect"

// leetcode-cn number: 5785. 合并若干三元组以形成目标三元组
func MergeTriplets(triplets [][]int, target []int) bool {
	match := func(r []int) bool {
		m := r[0] <= target[0] && r[1] <= target[1] && r[2] <= target[2]
		return m
	}

	max := func(a, b []int) []int {
		var r = make([]int, 3)
		for i := range a {
			if a[i] >= b[i] {
				r[i] = a[i]
			} else {
				r[i] = b[i]
			}
		}
		return r
	}

	var tFlag = false
	for i := range triplets {
		if !match(triplets[i]) {
			continue
		}

		if i+1 >= len(triplets) {
			break
		}

		if !match(triplets[i+1]) {
			triplets[i], triplets[i+1] = triplets[i+1], triplets[i]
			continue
		}

		if reflect.DeepEqual(triplets[i], target) || reflect.DeepEqual(triplets[i+1], target) {
			tFlag = true
			break
		}

		triplets[i+1] = max(triplets[i+1], triplets[i])

	}
	if !tFlag && reflect.DeepEqual(triplets[len(triplets)-1], target) {
		tFlag = true
	}
	return tFlag
}

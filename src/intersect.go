package src

func Intersect(nums1 []int, nums2 []int) []int {
	var longNums, shortNums []int
	if len(nums1) >= len(nums2) {
		longNums = nums1
		shortNums = nums2
	} else {
		longNums = nums2
		shortNums = nums1
	}
	var cases = make(map[int]int)
	for _, v := range shortNums {
		if _, ok := cases[v]; ok {
			cases[v] += 1
		} else {
			cases[v] = 1
		}
	}

	var r = make([]int, 0)
	for _, v := range longNums {
		if value, ok := cases[v]; ok && value != 0 {
			r = append(r, v)
			cases[v] -= 1
		}
	}
	return r
}

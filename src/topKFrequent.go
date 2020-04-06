package src

func TopKFrequent(nums []int, k int) []int {
	var cases = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		cases[nums[i]] += 1
	}

	var ans = make([]int, k)
	for i := 0; i < k; i++ {
		max := 0
		for k, v := range cases {
			if v > cases[max] {
				max = k
			}
		}

		ans[i] = max
		delete(cases, max)
	}
	return ans
}

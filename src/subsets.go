package src

func Subsets(nums []int) [][]int {
	var sets = make([][]int, 0)
	sets = append(sets, []int{})
	for i := 0; i < len(nums); i++ {
		t := make([]int, 0)
		for _, v := range sets {
			t = append([]int{}, v...)
			t = append(t, nums[i])
			sets = append(sets, t)
		}
	}
	return sets
}

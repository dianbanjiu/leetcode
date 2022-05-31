package perday

func TwoSum(nums []int, target int) []int {
	var leftMap = make(map[int]int)
	for i, num := range nums {
		var index int
		var ok bool
		if index, ok = leftMap[target-num]; ok {
			return []int{index, i}
		}
		leftMap[num] = i
	}
	return nil
}

package src

func Merge(intervals [][]int) [][]int {
	insertSort(intervals)
	return intervals
}

func insertSort(nums [][]int) {
	for i := 0; i < len(nums)-1; i++ {
		min := append([][]int{}, nums[i])
		for j := i + 1; j < len(nums); j++ {
			if nums[j][0] < min[0][0] {
				min[0] = nums[j]
			}
		}
		nums[i] = min[0]
	}
}

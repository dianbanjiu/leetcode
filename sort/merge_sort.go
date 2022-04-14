package sortCode

// 归并排序
func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	middle := len(nums) / 2
	left, right := mergeSort(nums[:middle]), mergeSort(nums[middle:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	var result = make([]int, 0)
	leftIndex, rightIndex := 0, 0
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] <= right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}

	if leftIndex < len(left) {
		result = append(result, left[leftIndex:]...)
	}

	if rightIndex < len(right) {
		result = append(result, right[rightIndex:]...)
	}

	return result
}

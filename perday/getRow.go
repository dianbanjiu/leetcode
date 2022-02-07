package perday

// leetcode-cn number: 119. 杨辉三角 II
// leetcode-cn url: https://leetcode-cn.com/problems/pascals-triangle-ii/
func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}

	var result = []int{1, 1}
	for i := 2; i <= rowIndex; i++ {
		var temp = make([]int, i+1)
		temp[0], temp[i] = 1, 1
		for j := 1; j < i; j++ {
			temp[j] = result[j-1] + result[j]
		}
		result = temp
	}
	return result
}

package perday

// 剑指 Offer 04. 二维数组中的查找
func FindNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	i, j := 0, len(matrix[0])-1

	for i < len(matrix) && j >= 0 {
		if matrix[i][j] == target {
			return true
		}

		if matrix[i][j] > target {
			j--
			continue
		}

		if matrix[i][j] < target {
			i++
		}
	}

	return false
}

package src

func FindNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	row, colum := 0, len(matrix[0])-1
	for row < len(matrix) && colum >= 0 {
		if matrix[row][colum] == target {
			return true
		} else if matrix[row][colum] > target {
			colum--
		} else {
			row++
		}
	}
	return false
}

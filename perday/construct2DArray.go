package perday

// construct2DArray leetcode-cn number: 5871. 将一维数组转变成二维数组
func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return nil
	}

	var result = make([][]int, m)
	var count = 0
	for i := 0; i < m; i++ {
		result[count] = original[i*n : (i+1)*n]
		count++
	}

	return result
}

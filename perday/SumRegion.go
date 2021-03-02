package perday

// leetcode-cn number: 304. 二维区域和检索 - 矩阵不可变
// url:https://leetcode-cn.com/problems/range-sum-query-2d-immutable/
type NumMatrix struct {
	Matrix [][]int
}

func MatrixConstructor(matrix [][]int) NumMatrix {
	return NumMatrix{matrix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	var sum int
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			sum += this.Matrix[i][j]
		}
	}
	return sum
}

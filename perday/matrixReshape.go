package perday

// leetcode-cn number: 566. 重塑矩阵
// url：https://leetcode-cn.com/problems/reshape-the-matrix/
// 解题思路：首先判断给出的行列数与原本的行列数是否匹配，匹配的情况下按照给出的行列数进行填充
func MatrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums)*len(nums[0]) != r*c {
		return nums
	}

	var result = make([][]int, r)
	numsr, numsc := 0, 0
	for ri := 0; ri < r; ri++ {
		result[ri] = make([]int, c)
		for ci := 0; ci < c; ci++ {
			result[ri][ci] = nums[numsr][numsc]
			numsc += 1
			if numsc >= len(nums[0]) {
				numsc = 0
				numsr += 1
			}
		}
	}

	return result
}

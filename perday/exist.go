package perday

// 剑指 Offer 12. 矩阵中的路径
func exist(board [][]byte, word string) bool {
	for i, bytes := range board {
		for j, _ := range bytes {
			if dfs(i, j, 0, word, board) {
				return true
			}
		}
	}
	return false
}

func dfs(i, j, k int, word string, board [][]byte) bool {
	if !((0 <= i && i < len(board)) && (0 <= j && j < len(board[i])) && word[k] == board[i][j]) {
		return false
	}
	if k == len(word)-1 {
		return true
	}
	board[i][j] = ' '
	res := dfs(i+1, j, k+1, word, board) || dfs(i-1, j, k+1, word, board) || dfs(i, j+1, k+1, word, board) || dfs(i, j-1, k+1, word, board)
	board[i][j] = word[k]
	return res
}

package perday

// 剑指 Offer 47. 礼物的最大价值
func MaxValue(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				grid[i][j] += grid[i][j-1]
				continue
			}
			if j == 0 {
				grid[i][j] += grid[i-1][j]
				continue
			}
			if grid[i][j-1] > grid[i-1][j] {
				grid[i][j] = grid[i][j] + grid[i][j-1]
			} else {
				grid[i][j] = grid[i][j] + grid[i-1][j]
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}

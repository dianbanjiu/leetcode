package perday

// 剑指 Offer 13. 机器人的运动范围
func movingCount(m int, n int, k int) int {
	var visited = make(map[[2]int]bool)

	return dfsMovingCount(0, 0, m, n, k, visited)
}

func dfsMovingCount(i, j, m, n, k int, visited map[[2]int]bool) int {
	if i < 0 || i >= m || j < 0 || j >= n || sumNumber(i)+sumNumber(j) > k || visited[[2]int{i, j}] {
		return 0
	}
	visited[[2]int{i, j}] = true
	return 1 + dfsMovingCount(i+1, j, m, n, k, visited) + dfsMovingCount(i, j+1, m, n, k, visited)
}

func sumNumber(n int) int {
	var sum int
	for n != 0 {
		temp := n % 10
		sum += temp
		n /= 10
	}
	return sum
}

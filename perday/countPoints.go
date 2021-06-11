package perday

// leetcode-cn number: 1828. 统计一个圆中点的数目
// leetcode url: https://leetcode-cn.com/problems/queries-on-number-of-points-inside-a-circle/

func CountPoints(points [][]int, queries [][]int) []int {
	var result = make([]int, len(queries))
	for i, query := range queries {
		r2 := query[2]*query[2]
		var count int
		for _, point := range points {
			baseFlag := (point[1]-query[1])*(point[1]-query[1])+(point[0]-query[0])*(point[0]-query[0])<=r2
			if baseFlag {
				count++
			}
		}
		result[i] = count
	}

	return result
}
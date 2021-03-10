package perday

// leetcode-cn number: LCP 17. 速算机器人
// url: https://leetcode-cn.com/problems/nGK0Fy/
func Calculate(s string) int {
	var x, y int
	for _, sb := range s {
		var ss = string(sb)
		if ss == "A" {
			x = 2*x + y
		} else {
			y = 2*y + x
		}
	}

	return x + y
}

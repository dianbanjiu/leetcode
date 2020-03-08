package src

func ClimbStairs(n int) int {
	var dplist = make([]int, n+1)
	for i := 0; i <= n; i++ {
		if i == 0 {
			dplist[i] = 0
		} else if i == 1 {
			dplist[i] = 1
		} else if i == 2 {
			dplist[i] = 2
		} else {
			dplist[i] = dplist[i-1] + dplist[i-2]
		}
	}
	return dplist[n]
}

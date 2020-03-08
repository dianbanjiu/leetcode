package src

func MySqrt(x int) int {
	var r int
	for i := 0; i <= x; i++ {
		if i*i <= x && (i+1)*(i+1) > x {
			r = i
			break
		}
	}
	return r
}

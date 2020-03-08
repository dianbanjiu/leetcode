package src

func FindContinuousSequence(target int) [][]int {
	var res = make([][]int, 0)
	var t = make([]int, 0)
	for i := 0; i <= target/2; i++ {
		t = append(t, i+1)
	}
	for i := 0; i < target/2; i++ {
		for j := i + 1; j <= target/2; j++ {
			if (i+j+2)*(j-i+1) == 2*target {
				res = append(res, t[i:j+1])
				break
			} else if (i+j+2)*(j-i+1) > 2*target {
				break
			}
		}
	}
	return res
}

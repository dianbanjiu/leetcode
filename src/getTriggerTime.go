package src

func GetTriggerTime(increase [][]int, requirements [][]int) []int {
	var ans = make([]int, len(requirements))
	for i := 0; i < len(requirements); i++ {
		ans[i] = -1
	}

	var curPower = make([]int, len(increase[0]))
	for i := 0; i < len(increase); i++ {
		for j := 0; j < len(requirements); j++ {
			for p := 0; p < len(increase[i]); p++ {
				if requirements[p][i] <= curPower[i] {
					ans[p] = i + 1
				}
			}
		}
	}
	return ans
}

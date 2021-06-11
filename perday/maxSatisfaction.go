package perday

import "sort"

// leetcode-cn number: 1402. 做菜顺序
// leetcode-cn url: https://leetcode-cn.com/problems/reducing-dishes/

func MaxSatisfaction(satisfaction []int) int {
	sort.Ints(satisfaction)
	var base int
	for i, s := range satisfaction {
		base += s * (i + 1)
	}

	for i := 1; i < len(satisfaction); i++ {
		var temp = satisfaction[i]
		var index = 2
		for j := i + 1; j < len(satisfaction); j++ {
			temp += satisfaction[j] * index
			index++
		}
		if temp < base {
			break
		}
		base = temp
	}

	if base < 0 {
		base = 0
	}
	return base
}

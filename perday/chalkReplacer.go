package perday

// leetcode-cn number: 5768. 找到需要补充粉笔的学生编号
func ChalkReplacer(chalk []int, k int) int {
	var sum int
	for _, c := range chalk {
		sum += c
	}
	for {
		if k-sum < 0 {
			break
		}
		k = k - sum
	}
	var last int
	for i := 0; ; i = (i + 1) % len(chalk) {
		k = k - chalk[i]
		if k < 0 {
			last = i
			break
		}

	}

	return last
}

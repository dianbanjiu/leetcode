package perday

import "strconv"

// 剑指 Offer 46. 把数字翻译成字符串
func TranslateNum(num int) int {
	numStr := strconv.Itoa(num)
	if len(numStr) == 1 {
		return 1
	}

	var temp [2]int
	value, _ := strconv.Atoi(numStr[:2])
	if value >= 0 && value < 26 {
		temp[0], temp[1] = 1, 2
	} else {
		temp[0], temp[1] = 1, 1
	}

	for i := 2; i < len(numStr); i++ {
		value, _ = strconv.Atoi(numStr[i-1 : i+1])
		if value >= 0 && value < 26 && numStr[i-1] != 48 {
			temp[0], temp[1] = temp[1], temp[0]+temp[1]
		} else {
			temp[0] = temp[1]
		}
	}

	return temp[1]
}

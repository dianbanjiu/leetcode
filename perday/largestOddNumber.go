package perday

import "strconv"

// leetcode-cn number: 5788. 字符串中的最大奇数
// url: https://leetcode-cn.com/contest/weekly-contest-246/problems/largest-odd-number-in-string/
func largestOddNumber(num string) string {
	var result string
	for i := len(num) - 1; i >= 0; i-- {
		n, _ := strconv.Atoi(string(num[i]))
		if n%2 == 1 {
			result = num[:i+1]
			break
		}
	}

	return result
}

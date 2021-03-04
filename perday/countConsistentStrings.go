package perday

import "strings"

// leetcode-cn number: 1684. 统计一致字符串的数目
// url: https://leetcode-cn.com/problems/count-the-number-of-consistent-strings/
func CountConsistentStrings(allowed string, words []string) int {
	var count = 0
	for _, word := range words {
		var flag bool
		for _, wc := range word {
			letter := string(wc)

			if !strings.Contains(allowed, letter) {
				flag = true
				break
			}
		}

		if !flag {
			count++
		}
	}

	return count
}

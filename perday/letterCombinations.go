package perday

import "strings"

// leetcode-cn number: 17. 电话号码的字母组合
// leetcode-cn url: https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
func letterCombinations(digits string) []string {
	var baseMap = map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	var result = make([]string, 0)
	for i := 0; i < len(digits); i++ {
		result = appendDigit(result, baseMap[digits[i]])
	}
	return result
}

func appendDigit(strs []string, str string) []string {
	var result = make([]string, 0)
	for i := 0; i < len(str); i++ {
		if len(strs) == 0 {
			return strings.Split(str, "")
		}
		stopLen := len(strs)
		for j := 0; j < stopLen; j++ {
			result = append(result, strs[j]+string(str[i]))
		}

	}
	return result
}

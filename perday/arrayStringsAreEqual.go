package perday

import "strings"

// leetcode-cn number: 1662. 检查两个字符串数组是否相等
// url: https://leetcode-cn.com/problems/check-if-two-string-arrays-are-equivalent/

func ArrayStringsAreEqual(word1 []string, word2 []string) bool {
	if strings.Join(word1, "") == strings.Join(word2, "") {
		return true
	}

	return false
}

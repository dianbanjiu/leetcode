package perday

import "strings"

// 剑指 Offer 58 - I. 翻转单词顺序
func ReverseWords(s string) string {
	if len(s) == 0 {
		return s
	}
	i := 0
	for ; i < len(s); i++ {
		if s[i] != ' ' {
			break
		}
	}

	j := len(s) - 1
	for ; j > i; j-- {
		if s[j] != ' ' {
			break
		}
	}

	s = s[i : j+1]
	i, j = len(s)-1, len(s)-1
	var result = make([]string, 0)
	for i >= 0 {
		if s[i] == ' ' && s[i+1] != ' ' {
			result = append(result, s[i+1:j+1])
		}
		if s[i] != ' ' && i+1 < len(s) && s[i+1] == ' ' {
			j = i
		}
		i--
	}
	result = append(result, s[i+1:j+1])
	return strings.Join(result, " ")
}

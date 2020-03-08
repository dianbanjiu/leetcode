package src

import "strings"

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	var letters = "0123456789qwertyuiopasdfghjklzxcvbnm"
	var (
		i    = 0
		j    = len(s) - 1
		flag = true
	)
	for i < len(s) && j >= 0 {
		if !strings.Contains(letters, string(s[i])) {
			i++
		} else if !strings.Contains(letters, string(s[j])) {
			j--
		} else if s[i] == s[j] {
			i++
			j--
		} else {
			flag = false
			break
		}
	}
	return flag
}

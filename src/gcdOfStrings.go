package src

import "strings"

func GcdOfStrings(str1 string, str2 string) string {
	var max int
	for i := 0; i < len(str1); i++ {
		if strings.Count(str1, str1[:i+1])*(i+1) == len(str1) && strings.Count(str2, str1[:i+1])*(i+1) == len(str2) {
			max = i+1
		}
	}
	return str1[:max]
}
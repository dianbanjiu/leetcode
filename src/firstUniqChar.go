package src

import "strings"

func FirstUniqChar(s string) int {
	var index = -1
	for i, v := range s {
		if strings.Count(s, string(v)) == 1 {
			index = i
			break
		}
	}
	return index
}

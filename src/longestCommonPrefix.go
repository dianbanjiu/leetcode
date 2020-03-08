package src

import "strings"

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var i int
	for ; i < len(strs[0]); i++ {
		flag := 0
		for j := 1; j < len(strs); j++ {
			if strings.Index(strs[j], strs[0][:i+1]) != 0 {
				flag = 1
				break
			}
		}
		if flag != 0 {
			break
		}
	}
	return strs[0][:i]
}

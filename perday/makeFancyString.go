package perday

import "strings"

// leetcode-cn number: 5193. 删除字符使字符串变好
func makeFancyString(s string) string {
	// 解法一，比较所有相邻的三个字符，时间复杂度 O(n)，空间复杂度 O(n)
	// 这种解法在 leetcode 会出现执行超时的情况
	//if len(s) < 3 {
	//	return s
	//}
	//
	//var result = s[:2]
	//for i := 2; i < len(s); i++ {
	//	if s[i] == result[len(result)-1] && s[i] == result[len(result)-2] {
	//		continue
	//	}
	//	result += string(s[i])
	//}
	//return result

	// 解法二：通过正则替换所有相邻的三个字符为两个，时间复杂度为 O(1)，空间复杂度为 O(1)
	var bases = "abcdefghijklmnopqrstuvwxyz"
	for _, base := range bases {
		ss := string(base)
		var flag bool
		for !flag {
			if !strings.Contains(s, ss+ss+ss) {
				flag = true
			} else {
				s = strings.ReplaceAll(s, ss+ss+ss, ss+ss)
			}
		}
	}
	return s
}

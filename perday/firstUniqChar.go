package perday

// 剑指 Offer 50. 第一个只出现一次的字符
func FirstUniqChar(s string) byte {
	var caseCount = [26]int{}
	for _, sc := range s {
		caseCount[sc-'a'] += 1
	}

	for _, sc := range s {
		if caseCount[sc-'a'] == 1 {
			return byte(sc)
		}
	}

	return ' '
}

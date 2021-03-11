package perday

// leetcode-cn number: 1704. 判断字符串的两半是否相似
// url: https://leetcode-cn.com/problems/determine-if-string-halves-are-alike/
func HalvesAreAlike(s string) bool {
	var baseMap = make(map[string]bool)
	for _, sc := range "aeiouAEIOU" {
		baseMap[string(sc)] = true
	}
	var frontCount, endCount int
	for i := 0; i < len(s)/2; i++ {
		if baseMap[string(s[i])] {
			frontCount++
		}
	}

	for i := len(s) / 2; i < len(s); i++ {
		if baseMap[string(s[i])] {
			endCount++
		}
	}
	return frontCount == endCount
}

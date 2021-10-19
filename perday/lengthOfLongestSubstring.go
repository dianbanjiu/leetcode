package perday

// 剑指 Offer 48. 最长不含重复字符的子字符串
func LengthOfLongestSubstring(s string) int {
	i := -1
	var temp int
	var dict = make(map[rune]int)
	for index, value := range s {
		if preIndex, ok := dict[value]; ok && i < preIndex {
			i = preIndex
		}
		if temp < index-i {
			temp = index - i
		}
		dict[value] = index
	}

	return temp
}

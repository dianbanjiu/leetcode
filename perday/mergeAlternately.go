package perday

// leetcode-cn number: 5685. 交替合并字符串
func MergeAlternately(word1 string, word2 string) string {
	var w1, w2 = 0, 0
	var result string
	for w1 < len(word1) && w2 < len(word2) {
		result += string(word1[w1]) + string(word2[w2])
		w1++
		w2++
	}

	result += word1[w1:] + word2[w2:]
	return result
}

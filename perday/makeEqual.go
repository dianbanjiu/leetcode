package perday

// leetcode-cn number: 5784. 重新分配字符使所有字符串都相等
func MakeEqual(words []string) bool {
	var wordsLen = len(words)
	var letters = make(map[rune]int)
	for _, word := range words {
		for _, w := range word {
			letters[w] += 1
		}
	}

	for _, v := range letters {
		if v%wordsLen != 0 {
			return false
		}
	}

	return true
}

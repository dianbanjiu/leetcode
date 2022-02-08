package perday

// leetcode-cn number: 1624. 两个相同字符之间的最长子字符串
// leetcode-cn url: https://leetcode-cn.com/problems/largest-substring-between-two-equal-characters/
// 此题题意中所说的两个相同字符之间可以包含本身，即 ababba 的正确答案为 babb
func maxLengthBetweenEqualCharacters(s string) int {
	type element struct {
		firstIndex int
		lastIndex  int
	}
	var result = -1
	var elementMap = make(map[byte]element)
	for i := 0; i < len(s); i++ {
		currentElement, ok := elementMap[s[i]]
		if !ok {
			currentElement.firstIndex = i
			elementMap[s[i]] = currentElement
			continue
		}
		currentElement.lastIndex = i
		elementMap[s[i]] = currentElement
		if currentElement.lastIndex-currentElement.firstIndex-1 > result {
			result = currentElement.lastIndex - currentElement.firstIndex - 1
		}
	}

	return result
}

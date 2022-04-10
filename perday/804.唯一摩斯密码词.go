package perday

// 题解：通过数组存储摩斯密码表，遍历传入的所有单词
// 通过一个临时变量存储其对应的摩斯密码
// 将计算出来的摩斯密码存储在一个 map 中
// 最终返回 map 的长度
// 时间复杂度：O(n*m)，n 是单词的数量，m 是其中每个单词的长度
// 空间复杂度：O(n)，需要一个 map 来存储每个单词的摩斯密码
func uniqueMorseRepresentations(words []string) int {
	var baseLetters = [26]string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	var result = make(map[string]struct{}, 0)
	for _, word := range words {
		var temp string
		for i := 0; i < len(word); i++ {
			temp += baseLetters[word[i]-'a']
		}
		result[temp] = struct{}{}
	}
	return len(result)
}

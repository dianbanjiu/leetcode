package perday

import "strings"

// leetcode-number: 5161. 可以输入的最大单词数
func canBeTypedWords(text string, brokenLetters string) int {
	var words = strings.Split(text, " ")
	var letters = strings.Split(brokenLetters, "")
	var count int
outloop:
	for _, word := range words {
		for _, letter := range letters {
			if strings.Contains(word, letter) {
				continue outloop
			}
		}
		count++
	}
	return count
}

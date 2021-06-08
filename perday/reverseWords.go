package perday

import "strings"

func ReverseWords(s string) string {
	var words string
	wordList := strings.Split(s, " ")
	for _, word := range wordList {
		var tempWord string
		for i := len(word) - 1; i >= 0; i-- {
			tempWord += string(word[i])
		}
		words += tempWord + " "
	}
	words = strings.TrimRight(words, " ")
	return words
}

package perday

// leetcode-cn number: 804. 唯一摩尔斯密码词
// url: https://leetcode-cn.com/problems/unique-morse-code-words/
func UniqueMorseRepresentations(words []string) int {
	var morseMap = map[string]string{
		"a": ".-",
		"b": "-...",
		"c": "-.-.",
		"d": "-..",
		"e": ".",
		"f": "..-.",
		"g": "--.",
		"h": "....",
		"i": "..",
		"j": ".---",
		"k": "-.-",
		"l": ".-..",
		"m": "--",
		"n": "-.",
		"o": "---",
		"p": ".--.",
		"q": "--.-",
		"r": ".-.",
		"s": "...",
		"t": "-",
		"u": "..-",
		"v": "...-",
		"w": ".--",
		"x": "-..-",
		"y": "-.--",
		"z": "--..",
	}
	var result = make(map[string]bool)
	for _, word := range words {
		var morseString string
		for _, w := range word {
			morseString += morseMap[string(w)]
		}
		result[morseString] = true
	}
	return len(result)
}

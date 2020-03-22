package src

func LongestPalindromelongestPalindrome(s string) int {
	var cases = make(map[string]int)
	for i := 0; i < len(s); i++ {
		c := string(s[i])
		if _, ok := cases[c]; ok {
			cases[c] += 1
		} else {
			cases[c] = 1
		}
	}

	var result int
	var maxOdd int
	for _, v := range cases {
		if v%2 == 0 {
			result += v
		} else if v-1 >= 2 {
			result += v - 1
			maxOdd = 1
		} else {
			maxOdd = 1
		}
	}
	return result + maxOdd
}

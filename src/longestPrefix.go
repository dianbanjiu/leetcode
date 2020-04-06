package src

func LongestPrefix(s string) string {
	var r string
	for i := 0; i < len(s)-1; i++ {
		if s[:i+1] == s[len(s)-i-1:] {
			r = s[:i+1]
		}
	}
	return r
}

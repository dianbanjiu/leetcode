package src

func ReverseLeftWords(s string, n int) string {
	return s[n%len(s):] + s[:n%len(s)]
}

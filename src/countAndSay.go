package src

import "strconv"

func CountAndSay(n int) string {
	var dplist = make([]string, n)
	dplist[0] = "1"

	for i := 1; i < n; i++ {
		temp := readString(dplist[i-1])
		dplist[i] = temp
	}
	return dplist[n-1]
}

func readString(s string) string {
	var c string
	var r string
	for i := 0; i < len(s); {
		if c == "" {
			c += string(s[i])
			i++
		} else if string(s[i]) == string(c[len(c)-1]) {
			c += string(s[i])
			i++
		} else {
			r += strconv.Itoa(len(c)) + string(c[len(c)-1])
			c = ""
		}

	}
	r += strconv.Itoa(len(c)) + string(c[len(c)-1])
	return r
}

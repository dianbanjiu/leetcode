package src

import "strconv"

func CompressString(S string) string {
	if len(S) == 0 {
		return S
	}
	var r string // result
	var count int
	var t = string(S[0])
	for i := 0; i < len(S); {
		if string(S[i]) == t {
			i++
			count++
		} else if count != 0 {
			r += t + strconv.Itoa(count)
			t = string(S[i])
			count = 0
		}
	}
	r += t + strconv.Itoa(count)
	if len(S) <= len(r) {
		return S
	}
	return r
}

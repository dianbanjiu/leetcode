package perday

// 541. 反转字符串 II
// leetcode-cn url: https://leetcode-cn.com/problems/reverse-string-ii/

func ReverseStr(s string, k int) string {
	var r string
	lastLen := len(s) % (2 * k)
	for i := 0; i < len(s)-lastLen; {
		r += reverse(s[i:i+k]) + s[i+k:i+2*k]
		i += 2 * k
	}

	if lastLen > 0 && lastLen < k {
		r += reverse(s[len(s)-lastLen:])
	} else if lastLen >= k && lastLen < 2*k {
		r += reverse(s[len(s)-lastLen:len(s)-lastLen+k]) + s[len(s)-lastLen+k:]
	}
	return r
}
func reverse(s string) string {
	var r string
	for i := len(s) - 1; i >= 0; i-- {
		r += string(s[i])
	}
	return r
}

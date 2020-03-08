package src

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var verify = make(map[string]int)
	for i := 0; i < len(s); i++ {
		if _, ok := verify[string(s[i])]; ok {
			verify[string(s[i])] += 1
		} else {
			verify[string(s[i])] = 1
		}
	}
	var flag = true
	for i := 0; i < len(t); i++ {
		if _, ok := verify[string(t[i])]; !ok {
			flag = false
			break
		}
	}
	return flag
}

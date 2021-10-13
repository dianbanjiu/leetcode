package perday

// 剑指 Offer 05. 替换空格
func replaceSpace(s string) string {
	var result = make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			result = append(result, []byte("%20")...)
			continue
		}
		result = append(result, s[i])
	}

	return string(result)
}

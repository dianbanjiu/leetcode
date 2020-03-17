package src

func CountCharacters(words []string, chars string) int {
	var cases = make(map[string]int)
	for i := 0; i < len(chars); i++ {
		if _, ok := cases[string(chars[i])]; ok {
			cases[string(chars[i])] += 1
		} else {
			cases[string(chars[i])] = 1
		}
	}

	var r int
	for i := 0; i < len(words); i++ {
		t := copyMap(cases)
		flag := 0
		for j := 0; j < len(words[i]); j++ {
			if _, ok := t[string(words[i][j])]; ok && t[string(words[i][j])] > 0 {
				t[string(words[i][j])] -= 1
			} else {
				flag = 1
				break
			}
		}
		if flag == 0 {
			r += len(words[i])
		}
	}
	return r
}

func copyMap(m map[string]int) map[string]int {
	var r = make(map[string]int)
	for k, v := range m {
		r[k] = v
	}
	return r
}

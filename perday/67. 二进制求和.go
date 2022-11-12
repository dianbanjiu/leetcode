package perday

func addBinary(a string, b string) string {
	a, b = fillZero(a, b)
	var temp int
	var result string
	i := len(a) - 1
	for i >= 0 {
		var at, bt int
		if a[i] == '1' {
			at = 1
		}
		if b[i] == '1' {
			bt = 1
		}

		tt := temp + at + bt
		if tt == 0 {
			result = "0" + result
			temp = 0
		} else if tt == 1 {
			result = "1" + result
			temp = 0
		} else if tt == 2 {
			result = "0" + result
			temp = 1
		} else {
			result = "1" + result
			temp = 1
		}

		i--
	}
	if temp == 1 {
		return "1" + result
	}
	return result
}

func fillZero(a, b string) (string, string) {
	lenA, lenB := len(a), len(b)
	if lenA > lenB {
		for i := 0; i < lenA-lenB; i++ {
			b = "0" + b
		}
	} else {
		for i := 0; i < lenB-lenA; i++ {
			a = "0" + a
		}
	}
	return a, b
}

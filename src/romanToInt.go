package src

func RomanToInt(s string) int {
	roman := map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
	}
	var (
		i   int
		sum int
	)
	for i < len(s) {
		if i < len(s)-1 {
			if v, ok := roman[s[i:i+2]]; ok {
				i = i + 2
				sum += v
				continue
			}
		}
		sum += roman[string(s[i])]
		i = i + 1
	}
	return sum
}

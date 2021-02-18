package perday

import (
	"math"
	"strings"
)

// leetcode-cn number: 8. 字符串转换整数（atoi）
// url: https://leetcode-cn.com/problems/string-to-integer-atoi/
func MyAtoi(s string) int {
	s = strings.TrimLeft(s," ")
	if len(s) == 0 {
		return 0
	}
	var pro bool
	var si int
	if string(s[0]) == "-" {
		pro = true
		si ++
	}else if string(s[0]) == "+"{
		si ++
	}

	var numbersMap = map[string]int{
		"0":0,
		"1":1,
		"2":2,
		"3":3,
		"4":4,
		"5":5,
		"6":6,
		"7":7,
		"8":8,
		"9":9,
	}
	var result int
	for ;si < len(s);si++ {
		sc := string(s[si])
		if !strings.Contains("0123456789", sc) {
			break
		}
		if pro && math.MinInt32 >= 0-(result*10+numbersMap[sc]) {
			result = 0-math.MinInt32
			break
		}else if math.MaxInt32 <= result*10+numbersMap[sc] {
			result = math.MaxInt32
			break
		}
		result = result *10 + numbersMap[sc]
	}
	if pro {
		result = 0-result
	}
	return result
}

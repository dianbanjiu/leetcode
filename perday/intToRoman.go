package perday

// leetcode-cn number: 12. 整数转罗马数字
// url: https://leetcode-cn.com/problems/integer-to-roman/

func IntToRoman(num int) string {
	var result string
	var romanList = "IVXLCDM"
	for i := 1; num != 0; num /= 10 {
		var s string
		var n = num%10
		switch n {
		case 4:
			s = string(romanList[i-1])+string(romanList[i])
		case 9:
			s = string(romanList[i-1])+string(romanList[i+1])
		default:
			if n>=5 {
				s += string(romanList[i])
				n -= 5
			}
			for j := 0; j < n; j++ {
				s += string(romanList[i-1])
			}

		}
		i += 2
		result = s + result
	}
	return result
}

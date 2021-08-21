package perday


// leetcode-cn number: 20. 有效的括号
// leetcode-cn url: https://leetcode-cn.com/problems/valid-parentheses/
func IsValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	var charStack = make([]string, 0)
	charStack = append(charStack, string(s[0]))
	for i := 1; i < len(s); i++ {
		temps := string(s[i])
		if len(charStack) == 0 {
			charStack = append(charStack, temps)
			continue
		}
		preChar := charStack[len(charStack)-1]
		ok := (temps == ")" && preChar == "(") || (temps == "]" && preChar == "[") || (temps == "}" && preChar == "{")
		if ok {
			charStack = charStack[:len(charStack)-1]
		} else {
			charStack = append(charStack, temps)
		}

	}

	return len(charStack) == 0

}

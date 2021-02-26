package perday

import "strings"

// leetcode-cn number: 1678. 设计 Goal 解析器
// url:https://leetcode-cn.com/problems/goal-parser-interpretation/
func Interpret(command string) string {
	var result string
	for i := 0; i < len(command); i++ {
		switch string(command[i]) {
		case "G":
			result += "G"
		case "(":
			if string(command[i+1]) == ")" {
				result += "o"
				i++
			} else {
				result += "al"
				i += 3
			}
		}
	}
	return result
}

func interpret(command string) string {
	command = strings.Replace(command, "()", "o", -1)
	command = strings.Replace(command, "(al)", "al", -1)

	return command
}

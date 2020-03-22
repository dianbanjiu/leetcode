package src

import (
	"strconv"
	"strings"
)

func EvalRPN(tokens []string) int {
	var data = make([]int, 0)
	var op = make([]string, 0)
	var opcase = "+-*/"
	var result int
	for i := 0; i < len(tokens);  {
		t := tokens[i]
		if !strings.Contains(opcase, t) {
			d, _ := strconv.Atoi(t)
			data = append(data, d)
			i++
			continue
		}
		switch compare(t, op[len(op)-1]) {
		case 0:
			result += calculate(data[len(data)-2], data[len(data)-1], op[len(op)-1])
			data = data[:len(data)-2]
			op = op[:len(op)-1]
		case 1:
			op = append(op, t)
			i++

		}
	}
}

package src

import (
	"math"
	"strings"
)

func TitleToNumber(s string) int {
	var letter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var r int
	for i := 0; i < len(s); i++ {
		r += int(math.Pow(26, float64(len(s)-i-1))) * (strings.Index(letter, string(s[i])) + 1)
	}
	return r
}

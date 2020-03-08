package leetcode

import (
	"fmt"
	"leetcode/src"
	"testing"
)

func TestTitleToNumber(t *testing.T) {
	in := "ZY"
	except := 1
	actual := src.TitleToNumber(in)
	if except != actual {
		fmt.Println("Actual got ", actual)
	}
}

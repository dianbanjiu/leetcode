package leetcode

import (
	"fmt"
	"leetcode/src"
	"testing"
)

func TestCQueue(t *testing.T){
	cq := src.Constructor()
	fmt.Println(cq.DeleteHead())
	cq.AppendTail(3)
	fmt.Println(cq.DeleteHead())
	fmt.Println(cq.DeleteHead())

}
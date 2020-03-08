package main

import (
	"fmt"
	"leetcode/src"
)

type intSlice []int

func main() {
	a := &src.ListNode{
		Val:  0,
		Next: nil,
	}
	var b = make(map[int]interface{})
	b[a.Val] = &a.Val
	fmt.Printf("%T", b[a.Val])
}

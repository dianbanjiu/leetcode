package leetcode

import (
	"fmt"
	"leetcode/src"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	l1 := &src.ListNode{
		Val:  2,
		Next: nil,
	}
	l2 := new(src.ListNode)

	actual := src.MergeTwoLists(l1.Next, l2.Next)
	for p := actual; p != nil; p = p.Next {
		fmt.Println(p.Val)
	}
}

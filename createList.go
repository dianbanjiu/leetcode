package main

import (
	"leetcode/src"
)

//从给定的数组创建单链表，并返回头节点
func CreateListFromArray(nums []int) *src.ListNode{
	if len(nums)==0{
		return nil
	}

	var head = new(src.ListNode)
	head.Val = nums[0]
	p := head
	for i:=1;i<len(nums);i++{
		t := &src.ListNode{
			Val: nums[i],
			Next:nil,
		}
		p.Next = t
		p = p.Next
	}
	return head
}
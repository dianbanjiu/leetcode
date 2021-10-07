package singleLinkedList

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

// createSingleLinkList 根据给定的数组创建单链表
func createSingleLinkedList(data []int) *Node {
	if len(data) == 0 {
		return nil
	}
	var root = new(Node)
	root.Val = data[0]
	temp := root
	for i := 1; i < len(data); i++ {
		temp.Next = &Node{Val: data[i]}
		temp = temp.Next
	}

	return root
}

// createSingleCircleLinkedList 根据给定的数组创建环状单链表
func createSingleCircleLinkedList(data []int) *Node {
	root := createSingleLinkedList(data)
	if root == nil {
		return nil
	}
	temp := root
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = root

	return root
}

// reverse 链表反转
func reverse(head *Node) *Node {
	if head == nil {
		return nil
	}

	next := head.Next
	head.Next = nil
	for next != nil {
		ptmp := next.Next
		next.Next = head
		head = next
		next = ptmp
	}

	return head
}

// printSingleLinkedList 打印链表
func printSingleLinkedList(head *Node) {
	next := head
	for next != nil {
		fmt.Println(next.Val)
		next = next.Next
	}
}

// circleCheck 检测环状链表
func circleCheck(head *Node) bool {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}

	slow := head
	fast := head.Next.Next
	for slow != nil && fast != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			fast = nil
		}
	}

	return false
}

// GetSingleLinkedListLen 获取链表长度
func GetSingleLinkedListLen(head *Node) int {
	var listLen int
	for head != nil {
		listLen++
		head = head.Next
	}

	return listLen
}

// RemoveNthNode 移除链表中倒数第 n 个节点
func RemoveNthNode(head *Node, n int) *Node {
	listLen := GetSingleLinkedListLen(head)
	// 排除异常情况
	if head == nil || listLen <= n || n < 0 {
		return head
	}

	// 头节点特殊处理
	if n == 0 {
		return head.Next
	}

	// 找到待删除节点的前一个节点
	next := head
	count := 1
	for next != nil {
		if count != n {
			next = next.Next
			count++
			continue
		}
		next.Next = next.Next.Next
		break

	}

	return head
}

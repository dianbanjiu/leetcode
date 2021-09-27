package singleLinkedList

type Node struct {
	Val int
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
	for i := 1; i< len(data); i++ {
		temp.Next = &Node{Val: data[i]}
		temp = temp.Next
	}

	return root
}
// createSingleCircleLinkedList 根据给定的数组创建环状单链表
func createSingleCircleLinkedList(data []int) *Node {
	root := createSingleLinkedList(data)
	if root==nil {
		return nil
	}
	temp := root
	for temp.Next !=nil {
		temp = temp.Next
	}
	temp.Next = root

	return root
}

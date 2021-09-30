package tree

type Node struct {
	Val int
	Left *Node
	Right *Node
}

// PreorderTraverse 前序遍历
func PreorderTraverse(root *Node) []int {
	if root == nil {
		return nil
	}

	var res = make([]int, 0)
	res = append(res, root.Val)
	res = append(res, PreorderTraverse(root.Left)...)
	res = append(res, PreorderTraverse(root.Right)...)
	return res
}

// PostOrderTraverse 后序遍历
func PostOrderTraverse(root *Node) []int {
	if root == nil {
		return nil
	}

	var res = make([]int, 0)
	res = append(res, PostOrderTraverse(root.Left)...)
	res = append(res, PostOrderTraverse(root.Right)...)
	res = append(res, root.Val)
	return res
}

// InOrderTraverse 中序遍历
func InOrderTraverse(root *Node) []int {
	if root == nil {
		return nil
	}

	var res = make([]int, 0)
	res = append(res, InOrderTraverse(root.Left)...)
	res = append(res, root.Val)
	res = append(res, InOrderTraverse(root.Right)...)
	return res
}



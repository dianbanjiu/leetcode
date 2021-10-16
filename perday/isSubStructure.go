package perday

// 剑指 Offer 26. 树的子结构
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	var currentList = make([]*TreeNode, 0)
	currentList = append(currentList, A)
	for len(currentList) != 0 {
		var nextList = make([]*TreeNode, 0)
		for _, node := range currentList {
			if recurSub(node, B) {
				return true
			}
			if node.Left != nil {
				nextList = append(nextList, node.Left)
			}

			if node.Right != nil {
				nextList = append(nextList, node.Right)
			}
		}
		currentList = nextList
	}
	return recurSub(A, B) || recurSub(A.Left, B) || recurSub(A.Right, B)
}

func recurSub(A, B *TreeNode) bool {
	if B == nil {
		return true
	}

	if A == nil || A.Val != B.Val {
		return false
	}

	return recurSub(A.Left, B.Left) && recurSub(A.Right, B.Right)
}

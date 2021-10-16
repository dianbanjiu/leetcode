package perday

// 剑指 Offer 27. 二叉树的镜像
func mirrorTree(root *TreeNode) *TreeNode {
	// way 1: 使用数组
	// if root == nil {
	// 	return nil
	// }

	// var currentList = make([]*TreeNode, 0)
	// currentList = append(currentList, root)
	// for len(currentList) != 0 {
	// 	var nextList = make([]*TreeNode, 0)
	// 	for _, node := range currentList {
	// 		if node.Left != nil {
	// 			nextList = append(nextList, node.Left)
	// 		}
	// 		if node.Right != nil {
	// 			nextList = append(nextList, node.Right)
	// 		}

	// 		node.Left, node.Right = node.Right, node.Left
	// 	}
	// 	currentList = nextList
	// }

	// return root

	// way 2: 递归
	if root == nil {
		return root
	}

	root.Right, root.Left = mirrorTree(root.Left), mirrorTree(root.Right)
	return root
}

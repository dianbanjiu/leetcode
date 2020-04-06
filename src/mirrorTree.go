package src

// 镜像树
func MirrorTree(root *TreeNode) *TreeNode {
	mirror(root)
	return root
}

// 递归地将所有节点的左右孩子互换即可
func mirror(root *TreeNode) {
	if root == nil || root.Left == nil && root.Right == nil {
		return
	}

	root.Left, root.Right = root.Right, root.Left
	mirror(root.Left)
	mirror(root.Right)
}

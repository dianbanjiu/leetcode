package src

// 判断一棵树是不是对称
func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var curLevelNode = make([]*TreeNode, 0)
	curLevelNode = append(curLevelNode, root)
	for len(curLevelNode) != 0 {
		for i := 0; i <= len(curLevelNode)/2; i++ {
			if curLevelNode[i] == nil && curLevelNode[len(curLevelNode)-1-i] != nil ||
				curLevelNode[i] != nil && curLevelNode[len(curLevelNode)-1-i] == nil ||
				curLevelNode[i] != nil && curLevelNode[len(curLevelNode)-1-i] != nil &&
					curLevelNode[i].Val != curLevelNode[len(curLevelNode)-i-1].Val {
				return false
			}
		}
		nextLevelNode := make([]*TreeNode, 0)
		for i := 0; i < len(curLevelNode); i++ {
			if curLevelNode[i] != nil {
				nextLevelNode = append(nextLevelNode, curLevelNode[i].Right, curLevelNode[i].Left)
			}
		}

		curLevelNode = nextLevelNode
	}
	return true
}

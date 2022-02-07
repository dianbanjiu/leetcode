package perday

// leetcode-cn number: 1614. 括号的最大嵌套深度
// url: https://leetcode-cn.com/problems/maximum-nesting-depth-of-the-parentheses/
func MaxDepth(s string) int {
	var max = 0
	var deep = 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" {
			deep += 1
			if deep > max {
				max = deep
			}
		} else if string(s[i]) == ")" {
			deep -= 1
		}
	}

	return deep
}

// 剑指 Offer 55 - I. 二叉树的深度
func maxDepth(root *TreeNode) int {
	return maxDepthDfs(root)
}

func maxDepthDfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

type NAnyTreeNode struct {
	Val      int
	Children []*NAnyTreeNode
}

// leetcode-cn number: 559. N 叉树的最大深度
// leetcode-cn url: https://leetcode-cn.com/problems/maximum-depth-of-n-ary-tree/
func nAryTreeMaxDepth(root *NAnyTreeNode) int {
	if root == nil {
		return 0
	}
	var nodes = make([]*NAnyTreeNode, 0)
	nodes = append(nodes, root)
	var count int
	for len(nodes) != 0 {
		var temp = make([]*NAnyTreeNode, 0)
		for _, node := range nodes {
			if node != nil {
				temp = append(temp, node.Children...)
			}
		}
		nodes = temp
	}
	return count
}

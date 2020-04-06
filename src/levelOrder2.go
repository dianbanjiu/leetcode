package src

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ans = make([][]int, 0)  //存放返回结果
	cur := make([]*TreeNode, 0) //存放当前层的节点
	cur = append(cur, root)     // 首先添加根节点

	// 遍历每层，循环中止条件是当层再无节点
	for len(cur) != 0 {
		t1 := make([]*TreeNode, 0)
		t2 := make([]int, len(cur))
		for i := 0; i < len(cur); i++ {
			t2[i] = cur[i].Val
			if cur[i].Left != nil {
				t1 = append(t1, cur[i].Left)
			}
			if cur[i].Right != nil {
				t1 = append(t1, cur[i].Right)
			}
		}
		ans = append(ans, t2)
		cur = t1
	}
	return ans
}

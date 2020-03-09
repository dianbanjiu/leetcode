package src

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var level = make([]*TreeNode, 0)
	level = append(level, root)
	var high = 1
	var flag = true
	// 层序遍历
	for !listIsNil(level) {
		temp := make([]*TreeNode, int(math.Pow(2, float64(high))))
		// 遍历获取下一层的数据
		for i := 0; i < len(level); i++ {
			if level[i] != nil && level[i].Left != nil {
				temp[2*i] = level[i].Left
			}
			if level[i] != nil && level[i].Right != nil {
				temp[2*i+1] = level[i].Right
			}
		}

		// 判断下一层的数据是否符合要求
		if !listIsNil(temp) && !levelJudge(temp) {
			flag = false
			break
		}
		level = temp
		high++
	}
	return flag
}

func listIsNil(t []*TreeNode) bool {
	var flag = true
	for _, v := range t {
		if v != nil {
			flag = false
			break
		}
	}
	return flag
}
func levelJudge(t []*TreeNode) bool {
	var flag = true
	for i := 0; i < len(t)/2; i++ {
		if (t[i] == nil && t[len(t)-1-i] != nil) || (t[i] != nil && t[len(t)-1-i] == nil) || ((t[i] != nil && t[len(t)-1-i] != nil) && (t[i].Val != t[len(t)-i-1].Val)) {
			flag = false
			break
		}
	}
	return flag

}

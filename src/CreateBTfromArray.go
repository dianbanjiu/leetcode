package src

import (
	"strconv"
)

// 由层序遍历的结果来创建一棵二叉树
func CreateBTfromArray1(nums []string) *TreeNode {
	// 1. 如果数组长度为 0, 说明树中无数据，返回 nil
	if len(nums) == 0 {
		return nil
	}

	// 2. 创建根节点，并为根节点赋值
	var root = new(TreeNode)
	if nums[0] == "nil" {
		root = nil
	} else {
		c, _ := strconv.Atoi(nums[0])
		root.Val = c
	}

	// 3. 创建一个数组保存当前层的节点
	var cur = make([]*TreeNode, 0)
	cur = append(cur, root)

	for i := 1; i < len(nums); {
		t := make([]*TreeNode, 0)
		// 4. 逐个取出数组中的节点，判断其左右孩子，
		// 若不为空则赋值，并将其添加到下一层节点的数组中
		for j := 0; j < len(cur); j++ {
			if nums[i] == "nil" && i%2 == 1 {
				cur[j].Left = nil
				i++
			} else if nums[i] != "nil" && i%2 == 1 {
				c, _ := strconv.Atoi(nums[i])
				cur[j].Left = &TreeNode{
					Val: c,
				}
				t = append(t, cur[j].Left)
				i++
			}
			if nums[i] == "nil" && i%2 == 0 {
				cur[j].Right = nil
				i++
			} else if nums[i] != "nil" && i%2 == 0 {
				c, _ := strconv.Atoi(nums[i])
				cur[j].Right = &TreeNode{
					Val: c,
				}
				t = append(t, cur[j].Right)
				i++
			}
		}
		cur = t
	}
	return root
}

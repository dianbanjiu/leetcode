package leetcode

import (
	"leetcode/src"
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	var node = new(src.TreeNode)
	node.Val = 1
	node.Left = &src.TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	node.Right = &src.TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	node.Left.Left = &src.TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	node.Left.Right = &src.TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	node.Right.Left = &src.TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	node.Right.Right = &src.TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	actual := src.IsSymmetric(node)
	if actual != true {
		t.Errorf("actual got %v", actual)
	}
}

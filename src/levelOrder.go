package src

func levelOrder(root *TreeNode) []int {
	var ans = make([]int, 0)
	var cur = make([]*TreeNode,0)
	cur = append(cur, root)
	for len(cur)!=0{
		t := make([]*TreeNode, 0)
		for i:=0;i<len(cur);i++{
			if cur[i]!=nil{
				ans = append(ans, cur[i].Val)
			}
			if cur[i].Left!=nil{
				t = append(t, cur[i].Left)
			}
			if cur[i].Right!=nil{
				t = append(t, cur[i].Right)
			}
		}
		cur = t
	}
	return ans
}
package perday

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
// 剑指 Offer 35. 复杂链表的复制
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	var listMap = make(map[*Node]*Node)
	var temp = head
	for temp != nil {
		listMap[temp] = &Node{Val: temp.Val}
		temp = temp.Next
	}

	temp = head
	for temp != nil {
		listMap[temp].Next = listMap[temp.Next]
		listMap[temp].Random = listMap[temp.Random]
		temp = temp.Next
	}
	return listMap[head]
}

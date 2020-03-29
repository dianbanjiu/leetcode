package src

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 暴力破解
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	var p *ListNode
	for pa := headA; pa != nil; pa = pa.Next {
		for pb := headB; pb != nil; pb = pb.Next {
			if pa == pb {
				p = pb
				break
			}
		}
		if p != nil {
			break
		}
	}
	return p
}

// 哈希表法
//func GetIntersectionNode(headA, headB *ListNode) *ListNode {
//	var nodeMap = make(map[*ListNode]int)
//	var pa = headA
//	for pa!=nil{
//		nodeMap[pa] = pa.Val
//	}
//	var pb = headB
//	for pb!=nil{
//		if _, ok := nodeMap[pb]; ok {
//			break
//		}
//	}
//	return pb
//}
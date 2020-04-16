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
func getIntersectionNode2(headA, headB *ListNode) *ListNode{
	var cases = make(map[interface{}]int)
	for p := headA; p != nil; p = p.Next {
		cases[p]=1
	}
	for p := headB; p != nil; p = p.Next {
		if _, ok := cases[p];ok {
			return p
		}
	}

	return nil
}

// 遍历
func getIntersectionNode3(headA, headB *ListNode) *ListNode {
	var la,lb int
	for p:=headA;p!=nil;p=p.Next{la++}
	for p:=headB;p!=nil;p=p.Next{lb++}

	pa := headA
	pb := headB
	if la > lb {
		for i := 0; i < la-lb; i++ {
			pa = pa.Next
		}
	}else {
		for i := 0; i < lb-la; i++ {
			pb = pb.Next
		}
	}

	for pa != pb {
		if pa != nil {
			pa = pa.Next
		}
		if pb != nil {
			pb=pb.Next
		}
	}
	return pa
}
package singleLinkedList

import "testing"

func TestFindMiddle(t *testing.T) {
	testCases := []struct {
		desc string
		in   []int
		out  int
	}{
		{
			desc: "空链表",
			in:   nil,
			out:  -1,
		},
		{
			desc: "单节点",
			in:   []int{1},
			out:  1,
		},
		{
			desc: "双节点",
			in:   []int{1, 2},
			out:  2,
		},
		{
			desc: "奇数多节点",
			in:   []int{1, 2, 3, 4, 5},
			out:  3,
		},
		{
			desc: "偶数多节点",
			in:   []int{1, 2, 3, 4, 5, 6, 7, 8},
			out:  5,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			realOut := FindMiddle(createSingleLinkedList(tC.in))
			if realOut != tC.out {
				t.Errorf("except %d, but got %d", tC.out, realOut)
			}
		})
	}
}

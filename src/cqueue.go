package src

// 用两个栈实现一个队列
type CQueue struct {
	Stack1 []int
	Stack2 []int
}

func Constructor() CQueue {
	var cq CQueue
	cq.Stack1 = make([]int,0)
	cq.Stack2 = make([]int, 0)
	return cq
}

func (this *CQueue) AppendTail(value int) {
	this.Stack1 = append(this.Stack1,value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.Stack1) == 0 {
		return -1
	}

	for i := len(this.Stack1)-1; i > 0; i-- {
		this.Stack2 = append(this.Stack2, this.Stack1[i])
		this.Stack1 = this.Stack1[:i]
	}
	ans := this.Stack1[0]
	this.Stack1 = this.Stack1[:0]
	for i := len(this.Stack2)-1; i >= 0; i-- {
		this.Stack1 = append(this.Stack1, this.Stack2[i])
		this.Stack2 = this.Stack2[:i]
	}
	return ans
}

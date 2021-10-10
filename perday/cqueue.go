package perday

// 剑指 Offer 09. 用两个栈实现队列
type CQueue struct {
	Data []int
}

func CQueueConstructor() CQueue {
	return CQueue{Data: make([]int, 0)}
}

func (this *CQueue) AppendTail(value int) {
	this.Data = append(this.Data, value)
}

func (this *CQueue) DeleteHead() int {
	var result = -1

	if len(this.Data) != 0 {
		result = this.Data[0]
		this.Data = this.Data[1:]
	}
	return result
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

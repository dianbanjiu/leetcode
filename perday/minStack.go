package perday

// 剑指 Offer 30. 包含min函数的栈
type MinStack struct {
	Data    []int
	MinData []int
}

/** initialize your data structure here. */
func MinStackConstructor() MinStack {
	return MinStack{Data: make([]int, 0)}
}

func (this *MinStack) Push(x int) {
	this.Data = append(this.Data, x)
	if len(this.MinData) == 0 || x <= this.MinData[len(this.MinData)-1] {
		this.MinData = append(this.MinData, x)
	}
}

func (this *MinStack) Pop() {
	x := this.Data[len(this.Data)-1]
	this.Data = this.Data[:len(this.Data)-1]
	if x == this.MinData[len(this.MinData)-1] {
		this.MinData = this.MinData[:len(this.MinData)-1]
	}
}

func (this *MinStack) Top() int {
	return this.Data[len(this.Data)-1]
}

func (this *MinStack) Min() int {
	return this.MinData[len(this.MinData)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

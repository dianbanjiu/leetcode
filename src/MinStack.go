package src

type MinStack struct {
	Data []int
}

/** initialize your data structure here. */
func StackConstructor() MinStack {
	var e MinStack
	e.Data = make([]int, 0)
	return e
}

func (this *MinStack) Push(x int) {
	this.Data = append(this.Data, x)
}

func (this *MinStack) Pop() {
	if len(this.Data) == 0 {
		return
	}
	this.Data = this.Data[:len(this.Data)-1]
}

func (this *MinStack) Top() int {
	if len(this.Data) == 0 {
		return -1
	}
	return this.Data[len(this.Data)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.Data) == 0 {
		return -1
	}
	var min = this.Data[0]
	for i := 1; i < len(this.Data); i++ {
		if min > this.Data[i] {
			min = this.Data[i]
		}
	}
	return min
}

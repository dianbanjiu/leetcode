package perday

// leetcode-cn number: 232. 用栈实现队列
// url: https://leetcode-cn.com/problems/implement-queue-using-stacks/
type MyQueue struct {
	Element []int
}

/** Initialize your data structure here. */
func QueueConstructor() MyQueue {
	return MyQueue{Element: []int{}}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.Element = append(this.Element, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.Empty() {
		return 0
	}

	var head int
	head = this.Element[0]
	if len(this.Element) == 1 {
		this.Element = []int{}
	} else {
		this.Element = this.Element[1:]
	}

	return head
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.Empty() {
		return 0
	}
	return this.Element[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.Element) == 0 {
		return true
	}

	return false
}

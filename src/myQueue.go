package src

type Stack struct {
	StackData []int
}
type MyQueue struct {
	QueueData []int
	QueueTemp []int
}


/** Initialize your data structure here. */
func MyQueueConstructor() MyQueue {
	var myQueue MyQueue
	return myQueue
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
    this.QueueData = append(this.QueueData, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
    if this.Empty() {
		return -1
	}

	for len(this.QueueData)>1{
		queueLength := len(this.QueueData)
		this.QueueTemp = append(this.QueueTemp,this.QueueData[queueLength-1])
		this.QueueData = this.QueueData[:queueLength-1]
	}
	queueHead := this.QueueData[0]
	this.QueueData = []int{}
    for len(this.QueueTemp)>0{
		tempQueueLength := len(this.QueueTemp)
		this.QueueData = append(this.QueueData, this.QueueTemp[tempQueueLength-1])
		this.QueueTemp = this.QueueTemp[:tempQueueLength-1]
	}

	return queueHead
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.Empty() {
		return -1
	}

	for len(this.QueueData)>1{
		queueLength := len(this.QueueData)
		this.QueueTemp = append(this.QueueTemp,this.QueueData[queueLength-1])
		this.QueueData = this.QueueData[:queueLength-1]
	}
	queueHead := this.QueueData[0]
    for len(this.QueueTemp)>0{
		tempQueueLength := len(this.QueueTemp)
		this.QueueData = append(this.QueueData, this.QueueTemp[tempQueueLength-1])
		this.QueueTemp = this.QueueTemp[:tempQueueLength-1]
	}

	return queueHead
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
   if len(this.QueueData)==0 {
	   return true
   }
   
   return false
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

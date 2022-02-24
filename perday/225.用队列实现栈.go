/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] 用队列实现栈
 */

// @lc code=start
package perday
type MyStack struct {
	Val []int
	Len int
}


func StackConstructor() MyStack {
	return MyStack{
		Val: []int{},
		Len: 0,
	}
}


func (this *MyStack) Push(x int)  {
	this.Len++
	this.Val = append(this.Val, x)
}


func (this *MyStack) Pop() int {
	top := this.Top()
	this.Len--
	this.Val = this.Val[:this.Len]
	return top
}


func (this *MyStack) Top() int {
	return this.Val[this.Len-1]
}


func (this *MyStack) Empty() bool {
	return len(this.Val) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end


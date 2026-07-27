package stackque

type MyQueue struct {
	Stack    []int
	OutStack []int
}

func Constructor() MyQueue {
	return MyQueue{
		Stack:    []int{},
		OutStack: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.Stack = append(this.Stack, x)
}

func (this *MyQueue) Pop() int {
	if len(this.OutStack) == 0 {
		for len(this.Stack) > 0 {
			this.OutStack = append(this.OutStack, this.Stack[len(this.Stack)-1])
			this.Stack = this.Stack[:len(this.Stack)-1]
		}
	}
	a := this.OutStack[len(this.OutStack)-1]
	this.OutStack = this.OutStack[:len(this.OutStack)-1]
	return a
}

func (this *MyQueue) Peek() int {
	if len(this.OutStack) == 0 {
		for len(this.Stack) > 0 {
			this.OutStack = append(this.OutStack, this.Stack[len(this.Stack)-1])
			this.Stack = this.Stack[:len(this.Stack)-1]
		}
	}
	a := this.OutStack[len(this.OutStack)-1]
	return a
}

func (this *MyQueue) Empty() bool {
	if len(this.Stack) > 0 || len(this.OutStack) > 0 {
		return false
	} else {
		return true
	}
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

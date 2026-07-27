package stackque

type MyStack struct {
	InnerQueue Queue
}

func NewMyStack() MyStack {
	t := []int{}
	a := Queue{
		inner: t,
	}
	return MyStack{
		InnerQueue: a,
	}
}

func (this *MyStack) Push(x int) {
	this.InnerQueue.push(x)
}

func (this *MyStack) Pop() int {
	size := this.InnerQueue.size()
	for i := size; i > 1; i-- {
		t, ok := this.InnerQueue.pop()
		if ok {
			this.InnerQueue.push(t)
		} else {
			break
		}
	}
	a, ok := this.InnerQueue.pop()
	if ok {
		return a
	} else {
		return 0
	}
}

func (this *MyStack) Top() int {
	size := this.InnerQueue.size()
	for i := size; i > 1; i-- {
		t, ok := this.InnerQueue.pop()
		if ok {
			this.InnerQueue.push(t)
		} else {
			break
		}
	}
	a, ok := this.InnerQueue.peek()
	t, ok1 := this.InnerQueue.pop()
	if ok1 {
		this.InnerQueue.push(t)
	} else {
		return 0
	}
	if ok {
		return a
	} else {
		return 0
	}
}

func (this *MyStack) Empty() bool {
	size := this.InnerQueue.size()
	if size == 0 {
		return true
	} else {
		return false
	}
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

type Queue struct {
	inner []int
}

func (this *Queue) push(x int) {
	this.inner = append(this.inner, x)
}

func (this *Queue) pop() (int, bool) {
	if len(this.inner) == 0 {
		return 0, false
	}
	t := this.inner[0]
	this.inner = this.inner[1:]
	return t, true
}

func (this *Queue) peek() (int, bool) {
	if len(this.inner) == 0 {
		return 0, false
	}
	t := this.inner[0]
	return t, true
}

func (this *Queue) size() int {
	return len(this.inner)
}

package list

type MyLinkedList struct {
	Next *MyLinkedList
	Val  int
}

func Constructor() MyLinkedList {
	//作为虚拟头结点
	return MyLinkedList{
		Next: nil,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 {
		return -1
	}
	cur := this.Next
	for i := 0; i < index; i++ {
		if cur == nil {
			return -1
		}
		cur = cur.Next
	}
	if cur == nil {
		return -1
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	cur := this
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = &MyLinkedList{
		Val:  val,
		Next: nil,
	}
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 {
		return
	}
	cur := this
	i := 0
	for ; i < index; i++ {
		if cur.Next == nil {
			return
		}
		cur = cur.Next
	}

	t := cur.Next
	cur.Next = &MyLinkedList{
		Next: t,
		Val:  val,
	}

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 {
		return
	}
	cur := this
	i := 0
	for ; i < index; i++ {
		if cur.Next == nil {
			return
		}
		cur = cur.Next
	}

	if cur.Next == nil {
		return
	}
	cur.Next = cur.Next.Next

}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

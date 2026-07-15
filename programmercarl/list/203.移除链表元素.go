package list

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{
		Val:  int(math.Pow(10, 4)) + 1,
		Next: head,
	}
	tempNode := dummy
	for tempNode.Next != nil {
		if tempNode.Next.Val == val {
			tempNode.Next = tempNode.Next.Next
		} else {
			tempNode = tempNode.Next
		}
	}
	return dummy.Next
}

package list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}
	tempNode := &ListNode{
		Next: head,
	}
	firstP := tempNode
	secondP := tempNode
	for i := 0; i < n; i++ {
		secondP = secondP.Next
	}
	for secondP.Next != nil {
		secondP = secondP.Next
		firstP = firstP.Next
	}
	firstP.Next = firstP.Next.Next
	return tempNode.Next
}

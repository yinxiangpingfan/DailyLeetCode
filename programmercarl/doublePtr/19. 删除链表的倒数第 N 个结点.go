package doubleptr

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	t := &ListNode{
		Next: head,
	}
	if head.Next == nil {
		return nil
	}
	firstP := t
	lastP := t
	for i := 0; i < n+1; i++ {
		lastP = lastP.Next
	}

	for lastP != nil {
		lastP = lastP.Next
		firstP = firstP.Next
	}
	firstP.Next = firstP.Next.Next
	return t.Next
}

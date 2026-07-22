package doubleptr

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	slowP := head
	fastP := head
	for fastP != nil && fastP.Next != nil {
		slowP = slowP.Next
		fastP = fastP.Next.Next
		if slowP == fastP {
			break
		}
	}
	if fastP == nil || fastP.Next == nil {
		return nil
	}
	for slowP != nil {
		if head == slowP {
			return slowP
		}
		head = head.Next
		slowP = slowP.Next
	}
	return nil
}

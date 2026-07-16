package list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slowP := head
	fastP := head
	for {
		if slowP.Next == nil || fastP == nil || fastP.Next == nil {
			return nil
		}
		slowP = slowP.Next
		fastP = fastP.Next.Next
		if slowP == fastP {
			break
		}
	}
	for {
		if head == slowP {
			break
		}
		head = head.Next
		slowP = slowP.Next
	}
	return head
}

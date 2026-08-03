package hot100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	t := headA
	c := 0
	d := headB
	for headA != nil {
		c++
		headA = headA.Next
	}
	for headB != nil {
		c--
		headB = headB.Next
	}
	if c < 0 {
		c = -c
		for i := 0; i < c; i++ {
			d = d.Next
		}
	} else {
		for i := 0; i < c; i++ {
			t = t.Next
		}
	}
	for d != nil {
		if t == d {
			return t
		}
		t = t.Next
		d = d.Next
	}
	return nil
}

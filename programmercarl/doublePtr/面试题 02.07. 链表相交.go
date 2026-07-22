package doubleptr

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA := 0
	lenB := 0
	tA := headA
	tB := headB
	for tA != nil {
		tA = tA.Next
		lenA++
	}
	for tB != nil {
		tB = tB.Next
		lenB++
	}
	n := 0
	tA = headA
	tB = headB
	if lenA > lenB {
		n = lenA - lenB
		for i := 0; i < n; i++ {
			tA = tA.Next
		}
	} else {
		n = lenB - lenA
		for i := 0; i < n; i++ {
			tB = tB.Next
		}
	}

	if tA == tB {
		return tA
	}
	for tA != nil {
		tA = tA.Next
		tB = tB.Next
		if tA == tB {
			return tA
		}
	}
	return nil
}

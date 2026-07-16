package list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	countA := 0
	countB := countA
	cur := headA
	for cur != nil {
		countA++
		cur = cur.Next
	}
	cur = headB
	for cur != nil {
		countB++
		cur = cur.Next
	}
	cur1 := cur
	if countA > countB {
		a := countA - countB
		cur = headA
		for i := 0; i < a; i++ {
			cur = cur.Next
		}
		cur1 = headB
	} else {
		a := countB - countA
		cur = headB
		for i := 0; i < a; i++ {
			cur = cur.Next
		}
		cur1 = headA
	}
	for cur1 != nil {
		if cur == cur1 {
			return cur1
		}
		cur = cur.Next
		cur1 = cur1.Next
	}
	return nil
}

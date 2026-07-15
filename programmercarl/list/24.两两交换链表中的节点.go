package list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	cur := head
	returnH := head.Next
	var pre *ListNode
	a := 0
	for cur != nil {
		if a == 0 {
			pre = cur
			cur = cur.Next
			a = 1
		} else {
			t := cur.Next
			cur.Next = pre
			cur = t
			if cur == nil {
				// 后面没有节点
				pre.Next = nil
			} else if cur.Next == nil {
				// 后面只剩一个节点，它不需要交换
				pre.Next = cur
			} else {
				// 后面至少有两个节点
				// cur.Next 是下一组交换后的头节点
				pre.Next = cur.Next
			}
			a = 0
		}
	}
	return returnH
}

package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	rs := make([]int, 0)
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node != nil {
			traversal(node.Left)

			traversal(node.Right)
			rs = append(rs, node.Val)
		}
	}
	traversal(root)
	return rs
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal2(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	rs := make([]int, 0)
	cur := root
	stack = append(stack, cur)
	for cur != nil || len(stack) > 0 {
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur != nil {
			rs = append(rs, cur.Val)
			stack = append(stack, cur.Left)
			stack = append(stack, cur.Right)
		}
	}
	for leftP, rightP := 0, len(rs)-1; leftP <= rightP; leftP, rightP = leftP+1, rightP-1 {
		t := rs[leftP]
		rs[leftP] = rs[rightP]
		rs[rightP] = t
	}
	return rs
}

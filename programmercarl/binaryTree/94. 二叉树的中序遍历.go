package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	rs := make([]int, 0)
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node != nil {
			traversal(node.Left)
			rs = append(rs, node.Val)
			traversal(node.Right)
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
func inorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	rs := make([]int, 0)
	cur := root
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			rs = append(rs, stack[len(stack)-1].Val)
			cur = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			cur = cur.Right
		}
	}
	return rs
}

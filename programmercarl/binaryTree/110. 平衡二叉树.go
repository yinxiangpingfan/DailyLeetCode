package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	rb := true
	var traversol func(node *TreeNode) int
	traversol = func(node *TreeNode) int {
		if !rb {
			return 0
		}
		if node == nil {
			return 0
		}
		l := traversol(node.Left)
		r := traversol(node.Right)
		if l-r > 1 || r-l > 1 {
			rb = false
			return 0
		}
		if l > r {
			return l + 1
		} else {
			return r + 1
		}
	}
	traversol(root)
	return rb
}

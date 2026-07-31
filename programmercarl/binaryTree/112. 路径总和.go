package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	count := 0
	a := false
	var traversol func(node *TreeNode)
	traversol = func(node *TreeNode) {
		if a {
			return
		}
		if node == nil {
			return
		}
		count += node.Val
		if node.Left == nil && node.Right == nil && count == targetSum {
			a = true
			return
		}
		traversol(node.Left)
		traversol(node.Right)
		count -= node.Val
	}
	traversol(root)
	return a
}

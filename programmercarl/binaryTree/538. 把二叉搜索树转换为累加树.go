package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	var traversol func(node *TreeNode)
	count := 0
	traversol = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversol(node.Right)
		a := node.Val
		node.Val += count
		count += a
		traversol(node.Left)
	}
	traversol(root)
	return root
}

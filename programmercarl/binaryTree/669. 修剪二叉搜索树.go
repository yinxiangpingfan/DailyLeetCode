package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	var traversol func(root *TreeNode) *TreeNode
	traversol = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		l := traversol(root.Left)
		r := traversol(root.Right)
		if root.Val > high {
			return l
		} else if root.Val < low {
			return r
		} else {
			root.Left = l
			root.Right = r
			return root
		}
	}
	return traversol(root)
}

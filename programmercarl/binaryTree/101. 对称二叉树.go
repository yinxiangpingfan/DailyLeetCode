package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	leftP := root
	rightP := root
	rb := true
	var traversol func(leftP *TreeNode, rightP *TreeNode)
	traversol = func(leftP *TreeNode, rightP *TreeNode) {
		if rb == false {
			return
		}
		if leftP != nil && rightP != nil {
			if leftP.Val == rightP.Val {
				traversol(leftP.Left, rightP.Right)
				traversol(leftP.Right, rightP.Left)
			} else {
				rb = false
				return
			}
		} else if leftP == nil && rightP == nil {
			return
		} else {
			rb = false
			return
		}
	}
	traversol(leftP, rightP)
	return rb
}

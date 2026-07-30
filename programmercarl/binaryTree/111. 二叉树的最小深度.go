package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := minDepth(root.Left)
	r := minDepth(root.Right)
	min := 0
	if l == 0 {
		return r + 1
	}
	if r == 0 {
		return l + 1
	}
	if l < r {
		min = l + 1
	} else {
		min = r + 1
	}
	return min
}

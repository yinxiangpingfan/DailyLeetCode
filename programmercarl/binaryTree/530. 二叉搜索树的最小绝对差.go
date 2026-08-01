package binarytree

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
	//中序遍历
	min := int(math.MaxInt64)
	var pre *TreeNode
	var traversol func(node *TreeNode)
	traversol = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversol(node.Left)
		if pre != nil {
			if node.Val-pre.Val < min {
				min = node.Val - pre.Val
			}
		}
		pre = node
		traversol(node.Right)
	}
	traversol(root)
	return min
}

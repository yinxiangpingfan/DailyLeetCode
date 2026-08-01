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
func isValidBST(root *TreeNode) bool {
	rb := true
	var traversol func(node *TreeNode, min, max int)
	traversol = func(node *TreeNode, min, max int) {
		if rb == false {
			return
		}
		if node == nil {
			return
		}
		if node.Val <= min || node.Val >= max {
			rb = false
			return
		}
		if node.Left == nil && node.Right == nil {

		} else if node.Left == nil && node.Right != nil {
			if !(node.Val < node.Right.Val) {
				rb = false
			}
			if node.Right.Val > max {
				max = node.Right.Val
			}
		} else if node.Left != nil && node.Right != nil {
			if !(node.Left.Val < node.Val && node.Right.Val > node.Val) {
				rb = false
			}
			if node.Left.Val < min {
				min = node.Left.Val
			}
			if node.Right.Val > max {
				max = node.Right.Val
			}
		} else {
			if !(node.Val > node.Left.Val) {
				rb = false
			}
			if node.Left.Val < min {
				min = node.Left.Val
			}
		}
		traversol(node.Left, min, node.Val)
		traversol(node.Right, node.Val, max)
	}
	traversol(root, -int(math.Pow(2, float64(31)))-1, int(math.Pow(2, float64(31))))
	return rb
}

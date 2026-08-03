package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var traversol func(node *TreeNode) *TreeNode
	traversol = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		l := traversol(node.Left)
		r := traversol(node.Right)
		if l != nil && r != nil {
			return node
		}
		if (node == p || node == q) && (l != nil || r != nil) {
			return node
		}
		if node == p || node == q {
			return node
		}
		if l != nil {
			return l
		} else if r != nil {
			return r
		}
		return nil
	}
	return traversol(root)
}

package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	count := 0
	var traversol func(node *TreeNode, s int)
	traversol = func(node *TreeNode, s int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			//叶子
			if s == 0 {
				count += node.Val
			}
			return
		}
		traversol(node.Left, 0)
		traversol(node.Right, 1)
	}
	traversol(root, 1)
	return count
}

package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	v := 0
	var l *TreeNode
	var r *TreeNode
	if root1 == nil {
		v = root2.Val
		l = mergeTrees(nil, root2.Left)
		r = mergeTrees(nil, root2.Right)
	} else if root2 == nil {
		v = root1.Val
		l = mergeTrees(root1.Left, nil)
		r = mergeTrees(root1.Right, nil)
	} else {
		v = root1.Val + root2.Val
		l = mergeTrees(root1.Left, root2.Left)
		r = mergeTrees(root1.Right, root2.Right)
	}
	return &TreeNode{
		Val:  v,
		Left: l, Right: r,
	}
}

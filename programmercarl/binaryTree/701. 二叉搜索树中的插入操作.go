package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{
			Val: val,
		}
	}
	var traversol func(node *TreeNode)
	var pre *TreeNode
	traversol = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversol(node.Left)
		if node.Val < val {
			pre = node
		} else {
			return
		}
		traversol(node.Right)
	}
	traversol(root)
	if pre == nil {
		t := root
		for t.Left != nil {
			t = t.Left
		}
		t.Left = &TreeNode{
			Val: val,
		}
		return root
	}
	t := pre.Right
	pre.Right = &TreeNode{
		Val:   val,
		Right: t,
	}
	return root
}

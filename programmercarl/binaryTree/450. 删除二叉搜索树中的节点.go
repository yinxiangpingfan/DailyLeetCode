package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	//五种情况
	if root == nil {
		return nil
	}
	//1
	if root.Val == key {
		if root.Left == nil {
			return root.Right
		} else {
			cur := root.Left
			t := cur
			for cur.Right != nil {
				cur = cur.Right
			}
			cur.Right = root.Right
			return t
		}
	}
	t := root
	var pre *TreeNode
	yes := false
	for t != nil {
		if key < t.Val {
			pre = t
			t = t.Left
		} else if key > t.Val {
			pre = t
			t = t.Right
		} else {
			yes = true
			break
		}
	}
	if !yes {
		return root
	}
	//2 3
	if pre.Left == t {
		if t.Right == nil {
			pre.Left = t.Left
		} else {
			pre.Left = t.Right
			cur := t.Right
			for cur.Left != nil {
				cur = cur.Left
			}
			cur.Left = t.Left
		}
	} else {
		if t.Left == nil {
			pre.Right = t.Right
		} else {
			pre.Right = t.Left
			cur := t.Left
			for cur.Right != nil {
				cur = cur.Right
			}
			cur.Right = t.Right
		}
	}
	return root
}

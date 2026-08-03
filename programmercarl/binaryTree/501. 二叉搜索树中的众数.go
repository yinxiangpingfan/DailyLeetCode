package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	max := -1
	var traversol func(node *TreeNode)
	var pre *TreeNode
	t := 0
	rb := make([]int, 0)
	traversol = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversol(node.Left)
		if pre == nil {
			t = 1
		} else {
			if node.Val == pre.Val {
				t++
			} else {
				if t > max {
					max = t
					rb = make([]int, 0)
					rb = append(rb, pre.Val)
				} else if t == max {
					rb = append(rb, pre.Val)
				}
				t = 1
			}
		}
		pre = node
		traversol(node.Right)
	}
	traversol(root)
	if t > max {
		max = t
		rb = make([]int, 0)
		rb = append(rb, pre.Val)
	} else if t == max {
		rb = append(rb, pre.Val)
	}
	return rb
}

package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal3(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	rs := make([]int, 0)
	if root != nil {
		stack = append(stack, root)
	} else {
		return nil
	}
	node := root
	for len(stack) > 0 {
		node = stack[len(stack)-1]
		if node != nil {
			stack = stack[:len(stack)-1]
			stack = append(stack, node)
			stack = append(stack, nil)
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		} else {
			stack = stack[:len(stack)-1]
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			rs = append(rs, node.Val)
		}
	}
	return rs
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal3(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	rs := make([]int, 0)
	if root != nil {
		stack = append(stack, root)
	} else {
		return nil
	}
	node := root
	for len(stack) > 0 {
		node = stack[len(stack)-1]
		if node != nil {
			stack = stack[:len(stack)-1]
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			stack = append(stack, node)
			stack = append(stack, nil)
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
		} else {
			stack = stack[:len(stack)-1]
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			rs = append(rs, node.Val)
		}
	}
	return rs
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal3(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	rs := make([]int, 0)
	if root != nil {
		stack = append(stack, root)
	} else {
		return nil
	}
	node := root
	for len(stack) > 0 {
		node = stack[len(stack)-1]
		if node != nil {
			stack = stack[:len(stack)-1]
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			stack = append(stack, node)
			stack = append(stack, nil)
		} else {
			stack = stack[:len(stack)-1]
			node = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			rs = append(rs, node.Val)
		}
	}
	return rs

}

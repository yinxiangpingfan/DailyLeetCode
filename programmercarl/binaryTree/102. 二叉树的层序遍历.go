package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	rs := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	max := 1
	t := 0
	ret := []int{}
	for i := 1; len(queue) > 0; i++ {
		a := queue[0]
		queue = queue[1:]
		if a.Left != nil {
			queue = append(queue, a.Left)
			t++
		}
		if a.Right != nil {
			queue = append(queue, a.Right)
			t++
		}
		ret = append(ret, a.Val)
		if i == max {
			rs = append(rs, ret)
			max += t
			t = 0
			ret = []int{}
		}
	}
	return rs
}

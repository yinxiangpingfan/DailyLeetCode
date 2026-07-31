package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
	//层序遍历
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 1)
	queue[0] = root
	maxN := 1
	rig := -1
	count := 0
	for i := 1; len(queue) > 0 && i <= maxN; i++ {
		v := queue[0]
		queue = queue[1:]
		if v.Right != nil {
			queue = append(queue, v.Right)
			count++
		}
		if v.Left != nil {
			queue = append(queue, v.Left)
			count++
		}
		if i == maxN {
			rig = v.Val
			maxN += count
			count = 0
		}
	}
	return rig
}

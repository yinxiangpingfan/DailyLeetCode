package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	max := -1
	cur := -1
	for k, v := range nums {
		if v > max {
			max = v
			cur = k
		}
	}
	leftN := nums[:cur]
	rightN := nums[cur+1:]
	left := constructMaximumBinaryTree(leftN)
	right := constructMaximumBinaryTree(rightN)
	return &TreeNode{
		Val:   nums[cur],
		Left:  left,
		Right: right,
	}
}

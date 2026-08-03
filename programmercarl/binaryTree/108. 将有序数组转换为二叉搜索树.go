package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	lens := len(nums)
	mid := lens / 2
	l := sortedArrayToBST(nums[0:mid])
	r := sortedArrayToBST(nums[mid+1:])
	return &TreeNode{
		Val:   nums[mid],
		Left:  l,
		Right: r,
	}
}

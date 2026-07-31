package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}
	cur := postorder[len(postorder)-1]
	count := 0
	for _, v := range inorder {
		if v == cur {
			break
		}
		count++
	}
	left := inorder[0:count]
	right := inorder[count+1:]
	left1 := postorder[0:count]
	right1 := postorder[count : len(postorder)-1]
	ln := buildTree(left, left1)
	rn := buildTree(right, right1)
	return &TreeNode{
		Val:   cur,
		Left:  ln,
		Right: rn,
	}
}

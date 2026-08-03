package binarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestorBST(root, p, q *TreeNode) *TreeNode {
	if q.Val < p.Val {
		q, p = p, q
	}
	for root != nil {
		if p.Val <= root.Val && root.Val <= q.Val {
			return root
		}
		if q.Val < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return nil
}

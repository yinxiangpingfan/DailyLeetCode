package binarytree

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	hL := getLeftHeight(root)
	hR := getRightHeight(root)

	if hL == hR {
		// 高度相同，说明当前完全二叉树是满二叉树
		return int(math.Pow(2, float64(hL))) - 1
	} else {
		// 高度不同时，递归统计左右子树
		left := countNodes(root.Left)
		right := countNodes(root.Right)
		return left + right + 1
	}
}

// 计算最左侧高度
func getLeftHeight(root *TreeNode) int {
	hL := 0

	for {
		if root == nil {
			break
		}
		hL++
		root = root.Left
	}

	return hL
}

// 计算最右侧高度
func getRightHeight(root *TreeNode) int {
	hR := 0

	for {
		if root == nil {
			break
		}
		hR++
		root = root.Right
	}

	return hR
}

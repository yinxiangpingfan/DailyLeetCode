package binarytree

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	rs := make([]string, 0)
	a := []byte{
		'-', '>',
	}
	path := make([]byte, 0)
	var traversol func(node *TreeNode)
	traversol = func(node *TreeNode) {
		if node == nil {
			return
		}
		v := []byte(strconv.Itoa(node.Val))
		path = append(path, v...)
		path = append(path, a...)
		if node.Left == nil && node.Right == nil {
			//为叶子节点
			rs = append(rs, string(path[:len(path)-2]))
			//回溯
			path = path[:len(path)-2]
			for len(path) > 0 && path[len(path)-1] != '>' {
				path = path[:len(path)-1]
			}
			return
		}
		traversol(node.Left)
		traversol(node.Right)
		path = path[:len(path)-2]
		for len(path) > 0 && path[len(path)-1] != '>' {
			path = path[:len(path)-1]
		}
	}
	traversol(root)
	return rs
}

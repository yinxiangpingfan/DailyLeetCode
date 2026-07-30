package binarytree

func preorderTraversal(root *TreeNode) []int {
	rs := make([]int, 0)
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node != nil {
			rs = append(rs, node.Val)
			traversal(node.Left)
			traversal(node.Right)
		}
	}
	traversal(root)
	return rs
}

func preorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	rs := make([]int, 0)
	stack = append(stack, root)
	for len(stack) >= 1 {
		a := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if a != nil {
			rs = append(rs, a.Val)
			stack = append(stack, a.Right)
			stack = append(stack, a.Left)
		}
	}
	return rs
}

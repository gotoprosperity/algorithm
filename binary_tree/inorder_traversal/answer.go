package inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	// return inorderTraversalRecursively(root)
	return inorderTraversalUnrecursionWithStack(root)
}

func inorderTraversalUnrecursionWithStack(n *TreeNode) []int {
	var (
		s    = []*TreeNode{}
		ret  = []int{}
		node = n
	)
	for node != nil || len(s) > 0 {
		for node != nil {
			s = append(s, node)
			node = node.Left
		}
		if len(s) > 0 {
			node = s[len(s)-1]
			s = s[:len(s)-1]
			ret = append(ret, node.Val)
			node = node.Right
		}
	}
	return ret
}

func inorderTraversalRecursively(n *TreeNode) []int {
	if n == nil {
		return []int{}
	}
	var ret []int
	if n.Left != nil {
		ret = inorderTraversalRecursively(n.Left)
	}
	ret = append(ret, n.Val)
	if n.Right != nil {
		ret = append(ret, inorderTraversalRecursively(n.Right)...)
	}
	return ret
}

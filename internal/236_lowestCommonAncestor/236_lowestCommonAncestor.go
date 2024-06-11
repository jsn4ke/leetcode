package leetcode_236_lowestCommonAncestor

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q {
		return root
	}
	if nil == root {
		return nil
	}
	a := lowestCommonAncestor(root.Left, p, q)
	b := lowestCommonAncestor(root.Right, p, q)
	if nil == a {
		return b
	}
	if nil == b {
		return a
	}
	return root
}

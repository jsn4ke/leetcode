package leetcode_543_diameterOfBinaryTree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var (
		traverse func(node *TreeNode) int
		ans      int
	)
	traverse = func(node *TreeNode) int {
		if nil == node {
			return 0
		}
		leftDepth := traverse(node.Left)
		rightDepth := traverse(node.Right)
		ans = max(ans, leftDepth+rightDepth)
		return 1 + max(leftDepth, rightDepth)
	}
	traverse(root)
	return ans
}

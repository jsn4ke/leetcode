package leetcode_98_isValidBST

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if nil == root {
		return true
	}
	var (
		dfs func(node *TreeNode, minVal, maxVal int) bool
	)
	dfs = func(node *TreeNode, minVal, maxVal int) bool {
		if nil == node {
			return true
		}
		if node.Val <= minVal || node.Val >= maxVal {
			return false
		}
		return dfs(node.Left, minVal, node.Val) && dfs(node.Right, node.Val, maxVal)

	}
	return dfs(root, math.MinInt32-1, math.MaxInt32+1)
}

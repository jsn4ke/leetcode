package leetcode_124_maxPathSum

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	var (
		ans int = -3_0000_0000
	)
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if nil == node {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		ret := max(node.Val, max(node.Val+left, node.Val+right))
		ans = max(ans, max(ret, node.Val+left+right))
		return ret
	}
	dfs(root)
	return ans
}

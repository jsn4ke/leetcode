package leetcode_230_kthSmallest

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var dfs func(in *TreeNode)
	var find bool
	var ans int
	dfs = func(in *TreeNode) {
		if find || nil == in {
			return
		}
		dfs(in.Left)
		k--
		if 0 == k {
			ans = in.Val
			find = true
			return
		}
		dfs(in.Right)
	}
	dfs(root)
	return ans
}

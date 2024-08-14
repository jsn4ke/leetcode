package leetcode_652_findDuplicateSubtrees

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var (
		ans   []*TreeNode
		cache = make(map[string]int)
	)
	var dfs func(node *TreeNode) string
	dfs = func(node *TreeNode) string {
		if node == nil {
			return "."
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		subtree := fmt.Sprintf("%d,%s,%s", node.Val, left, right)
		if hit := cache[subtree]; 1 == hit {
			ans = append(ans, node)
		}
		cache[subtree]++
		return subtree
	}

	dfs(root)
	return ans
}

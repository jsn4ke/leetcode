package leetcode_437_pathSum

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	var ans int

	var stack []int

	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if nil == node {
			return
		}
		if node.Val == targetSum {
			ans++
		}
		for i := range stack {
			v := stack[i] + node.Val
			if v == targetSum {
				ans++
			}
			stack[i] = v
		}
		stack = append(stack, node.Val)
		dfs(node.Left)

		dfs(node.Right)
		stack = stack[:len(stack)-1]
		for i := range stack {
			stack[i] -= node.Val
		}
	}

	dfs(root)
	return ans
}

package leetcode_114_flatten

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	var dfs func(node *TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if nil == node {
			return nil
		}
		var (
			left  = node.Left
			right = node.Right
			tail1 *TreeNode
			tail2 *TreeNode
		)
		node.Left = nil
		node.Right = nil
		if nil != left {
			tail1 = dfs(left)
			node.Right = left
			node = tail1
		}
		if nil != right {
			tail2 = dfs(right)
			node.Right = right
			node = tail2
		}
		return node
	}
	dfs(root)

}

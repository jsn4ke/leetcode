package leetcode_108_sortedArrayToBST

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	var dfs func(l, r int) *TreeNode
	dfs = func(l, r int) *TreeNode {
		if l >= r {
			return nil
		}
		mid := (l + r) / 2
		node := new(TreeNode)
		node.Val = nums[mid]
		node.Left = dfs(l, mid)
		node.Right = dfs(mid+1, r)
		return node
	}
	return dfs(0, len(nums))
}

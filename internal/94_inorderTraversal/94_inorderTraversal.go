package leetcode_94_inorderTraversal

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if nil == root {
		return nil
	}
	left := inorderTraversal(root.Left)
	left = append(left, root.Val)
	right := inorderTraversal(root.Right)
	return append(left, right...)
}

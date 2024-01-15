package leetcode_199_rightSideView

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var (
		heap []*TreeNode
	)
	heap = append(heap, root)
}

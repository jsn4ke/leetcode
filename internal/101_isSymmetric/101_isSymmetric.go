package leetcode_101_isSymmetric

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if nil == root {
		return true
	}
	var isSame func(left, right *TreeNode) bool
	isSame = func(left, right *TreeNode) bool {
		if nil == left && nil != right {
			return false
		}
		if nil != left && nil == right {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		if !isSame(left.Left, right.Right) {
			return false
		}
		if !isSame(left.Right, right.Left) {
			return false
		}
		return true
	}
	return isSame(root.Left, root.Right)
}

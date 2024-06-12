package leetcode_102_levelOrder

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var (
		ans [][]int
		sk  []*TreeNode
	)
	if nil == root {
		return nil
	}
	sk = append(sk, root)
	for 0 != len(sk) {
		var tmp []*TreeNode
		var last []int
		for _, v := range sk {
			last = append(last, v.Val)
			if nil != v.Left {
				tmp = append(tmp, v.Left)
			}
			if nil != v.Right {
				tmp = append(tmp, v.Right)
			}
		}
		ans = append(ans, last)
		sk = tmp
	}
	return ans
}

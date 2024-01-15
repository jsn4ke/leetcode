package leetcode_199_rightSideView

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if nil == root {
		return nil
	}
	var (
		heap []*TreeNode
	)
	heap = append(heap, root)
	for 0 != len(heap) {
		cur := heap[0]
		if nil != cur.Left {
			heap = append(heap, cur.Left)
		}
		if nil != cur.Right {
			heap = append(heap, cur.Right)
		}
		heap = heap[1:]
	}
}

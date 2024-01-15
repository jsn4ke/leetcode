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
	type nodeWithLayer struct {
		*TreeNode
		level int
	}
	var (
		heap      []*nodeWithLayer
		ret       []int
		lastLayer = -1
	)
	heap = append(heap, &nodeWithLayer{
		TreeNode: root,
		level:    0,
	})
	for 0 != len(heap) {
		cur := heap[0]
		if cur.level != lastLayer {
			ret[cur.level] = cur.Val
		}
		if nil != cur.Left {
			heap = append(heap, &nodeWithLayer{
				TreeNode: cur.Left,
				level:    cur.level + 1,
			})
		}
		if nil != cur.Right {
			heap = append(heap, cur.Right)
		}
		heap = heap[1:]
	}
	return
}

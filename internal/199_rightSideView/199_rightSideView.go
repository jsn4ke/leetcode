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
		heap []*nodeWithLayer
		ret  []int
	)
	heap = append(heap, &nodeWithLayer{
		TreeNode: root,
		level:    0,
	})
	for 0 != len(heap) {
		cur := heap[0]
		if cur.level >= len(ret) {
			ret = append(ret, cur.Val)
		}
		ret[cur.level] = cur.Val
		if nil != cur.Left {
			heap = append(heap, &nodeWithLayer{
				TreeNode: cur.Left,
				level:    cur.level + 1,
			})
		}
		if nil != cur.Right {
			heap = append(heap, &nodeWithLayer{
				TreeNode: cur.Right,
				level:    cur.level + 1,
			})
		}
		heap = heap[1:]
	}
	return ret
}

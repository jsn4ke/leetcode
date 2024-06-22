package leetcode_272_closestKValues

import "container/heap"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	_ heap.Interface = (*hp)(nil)
)

type pair struct {
	a float64
	b int
}

type hp []pair

// Len implements heap.Interface.
func (h hp) Len() int {
	return len(h)
}

// Less implements heap.Interface.
func (h hp) Less(i int, j int) bool {
	return h[i].a < h[j].a
}

// Pop implements heap.Interface.
func (h *hp) Pop() any {
	last := h.Len() - 1
	v := (*h)[last]
	*h = (*h)[:last]
	return v
}

// Push implements heap.Interface.
func (h *hp) Push(x any) {
	*h = append((*h), x.(pair))
}

// Swap implements heap.Interface.
func (h hp) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func closestKValues(root *TreeNode, target float64, k int) []int {
	var arr hp
	abs := func(v float64) float64 {
		if v-target < 0 {
			return target - v
		} else {
			return v - target
		}
	}
	var insert = func(in *TreeNode) {
		heap.Push(&arr, pair{abs(float64(in.Val)), in.Val})
	}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if nil == node {
			return
		}
		insert(node)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	var ans []int
	for i := 0; i < k; i++ {
		ans = append(ans, heap.Pop(&arr).(pair).b)
	}
	return ans
}

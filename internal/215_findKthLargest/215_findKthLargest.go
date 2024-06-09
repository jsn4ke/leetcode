package leetcode_215_findKthLargest

import "container/heap"

var (
	_ heap.Interface = (*small)(nil)
	_ heap.Interface = (*big)(nil)
)

type small []int

// Len implements heap.Interface.
func (s small) Len() int {
	return len(s)
}

// Less implements heap.Interface.
func (s small) Less(i int, j int) bool {
	return s[i] < s[j]
}

// Pop implements heap.Interface.
func (s *small) Pop() any {
	s1 := *s
	idx := len(s1) - 1
	ret := s1[idx]
	*s = s1[:idx]
	return ret
}

// Push implements heap.Interface.
func (s *small) Push(x any) {
	s1 := *s
	s1 = append(s1, x.(int))
	*s = s1
}

// Swap implements heap.Interface.
func (s small) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}

func findKthLargest(nums []int, k int) int {
	var (
		sa small
	)
	for _, v := range nums {
		if sa.Len() < k {
			heap.Push(&sa, v)
		} else {
			if sa[0] < v {
				heap.Push(&sa, v)
				heap.Pop(&sa)
			}
		}
	}
	return sa[0]
}

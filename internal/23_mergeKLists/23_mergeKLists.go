package leetcode_23_mergeKLists

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var (
		n = len(lists)
	)
	if 0 == n {
		return nil
	}
	var merge func(l, r int) *ListNode
	merge = func(l, r int) *ListNode {
		if r >= n {
			return lists[l]
		}
		var (
			dummy = new(ListNode)
			a, b  = lists[l], lists[r]
		)
		p := dummy
		for nil != a && nil != b {
			if a.Val > b.Val {
				a, b = b, a
			}
			p.Next = a
			p = p.Next
			a = a.Next
		}
		if nil == a {
			p.Next = b
		} else {
			p.Next = a
		}
		return dummy.Next

	}
	for step := 1; step <= n; step <<= 1 {
		for i := 0; i < n; i += step << 1 {
			lists[i] = merge(i, i+step)
		}
	}
	return lists[0]
}

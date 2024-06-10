package leetcode_148_sortList

import (
	"sort"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	var (
		arr []int
		p   = head
	)
	for nil != p {
		arr = append(arr, p.Val)
		p = p.Next
	}
	sort.Ints(arr)
	p = head
	for _, v := range arr {
		p.Val = v
		p = p.Next
	}
	return head
}

func sortList2(head *ListNode) *ListNode {
	var (
		n     int
		dummy = new(ListNode)
	)
	dummy.Next = head
	for nil != head {
		n++
		head = head.Next
	}
	var merge func(left, right *ListNode) *ListNode
	merge = func(left, right *ListNode) *ListNode {
		dummy := new(ListNode)
		p := dummy
		for nil != left && nil != right {
			if left.Val > right.Val {
				left, right = right, left
			}
			p.Next = left
			left = left.Next
			p = p.Next
		}
		if nil == left {
			left = right
		}
		p.Next = left
		return dummy.Next
	}
	for step := 1; step < n; step <<= 1 {
		pre := dummy
		cur := dummy.Next
		for nil != cur {

			var (
				head1 *ListNode
				head2 *ListNode
				next  *ListNode
			)

			head1 = cur
			for i := 1; i < step && nil != cur.Next; i++ {
				cur = cur.Next
			}
			head2 = cur.Next
			cur.Next = nil
			cur = head2
			if nil != cur {
				for i := 1; i < step && nil != cur.Next; i++ {
					cur = cur.Next
				}
			}
			if nil != cur {
				next = cur.Next
				cur.Next = nil
			}
			pre.Next = merge(head1, head2)
			for nil != pre.Next {
				pre = pre.Next
			}
			cur = next
		}
	}
	return dummy.Next
}

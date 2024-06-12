package leetcode_24_swapPairs

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	var (
		dummy = new(ListNode)
		prev  = dummy
		cur   = head
	)
	dummy.Next = head
	for nil != cur && nil != cur.Next {
		first := cur
		second := cur.Next
		next := cur.Next.Next
		cur.Next.Next = nil
		second.Next = first
		prev.Next = second
		prev = first
		cur = next
		first.Next = next
	}
	return dummy.Next
}

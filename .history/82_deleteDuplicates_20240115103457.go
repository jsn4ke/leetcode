package leatcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	var (
		dummy = &ListNode{}
		pre   = dummy
		cur   = head
	)
	dummy.Next = head
}
